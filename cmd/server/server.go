package server

import (
	"embed"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/amarps/joebot/src/models"
	"github.com/amarps/joebot/src/server"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cobra"
)

type ServerCommand struct {
	Command *cobra.Command
	Config  *Server
}

type Server struct {
	Port          int
	WebPortalPort int
}

//go:embed web_portal_html/*
var webPortalAssets embed.FS

func Init() *ServerCommand {
	res := ServerCommand{
		Config: &Server{},
		Command: &cobra.Command{
			Use:   "server",
			Short: "Server Mode",
		},
	}

	res.Command.PersistentFlags().IntVarP(&res.Config.Port, "port", "p", 13579, "Port For Listening Slave Machine, Default = 13579")
	res.Command.PersistentFlags().IntVarP(&res.Config.WebPortalPort, "web-portal-port", "w", 8080, "Port For The Web Portal, Default = 8080")

	res.Command.Run = func(cmd *cobra.Command, args []string) {
		s := server.NewServer(nil)
		s.Start(res.Config.Port)

		e := echo.New()
		v1 := e.Group("/api")

		// if username != nil && password != nil && *username != "" && *password != "" {
		// 	v1.Use(middleware.BasicAuth(func(user, pw string, c echo.Context) (bool, error) {
		// 		if user == *username && pw == *password {
		// 			return true, nil
		// 		}
		// 		return false, nil
		// 	}))
		// }
		webPortalAssetsFS := WebPortalAssetsFS()

		v1.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		}))
		e.GET("/", func(c echo.Context) error {
			f, err := webPortalAssetsFS.Open("index.html")
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			b, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatal(err)
			}
			return c.HTML(200, string(b))
		})
		// e.GET("/*", echo.WrapHandler(joebot_html.Handler))
		e.GET("/*", echo.WrapHandler(http.FileServer(http.FS(webPortalAssetsFS))))
		v1.GET("/clients", func(c echo.Context) error {
			return c.JSON(http.StatusOK, s.GetClientsList())
		})
		v1.POST("/client/:id", func(c echo.Context) error {
			type msg struct {
				Message string `json:"message"`
			}

			client, err := s.GetClientById(c.Param("id"))
			if err != nil {
				return c.JSON(http.StatusNotFound, msg{err.Error()})
			}
			portStr := c.FormValue("target_client_port")
			port, err := strconv.Atoi(portStr)
			if err != nil || port <= 0 {
				return c.JSON(http.StatusBadRequest, msg{"Invalid target_client_port"})
			}

			portTunnelInfo, err := client.CreateTunnel(port)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, msg{err.Error()})
			}

			return c.JSON(http.StatusOK, portTunnelInfo)
		})
		v1.POST("/bulk-install", func(c echo.Context) error {
			json := models.BulkInstallInfo{}

			if err := c.Bind(&json); err != nil {
				return err
			}
			result, err := s.BulkInstallJoebot(json)
			if err != nil {
				return err
			}

			return c.String(http.StatusOK, result)
		})
		e.Start(":" + strconv.Itoa(res.Config.WebPortalPort))
	}

	return &res
}

func WebPortalAssetsFS() fs.FS {
	assetsFs, _ := fs.Sub(webPortalAssets, "web_portal_html")
	return assetsFs
}
