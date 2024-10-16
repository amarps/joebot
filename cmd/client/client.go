package client

import (
	"sync"

	"github.com/amarps/joebot/src/client"
	"github.com/spf13/cobra"
)

type ClientCommand struct {
	Command *cobra.Command
	Config  *ClientConfig
}

type ClientConfig struct {
	ServerIp              string
	ServerPort            int
	AllowedPortLowerBound int
	AllowedPortUpperBound int
	Tags                  []string
	FileBrowserDir        string
}

func Init() *ClientCommand {
	res := ClientCommand{
		Config: &ClientConfig{},
		Command: &cobra.Command{
			Use:   "client",
			Short: "Client Mode",
		},
	}

	res.Command.PersistentFlags().StringVar(&res.Config.ServerIp, "ip", "", "Server IP")
	res.Command.PersistentFlags().IntVarP(&res.Config.ServerPort, "port", "p", 13579, "Server Port, Default=13579")
	res.Command.PersistentFlags().IntVarP(&res.Config.AllowedPortLowerBound, "allowed-port-lower-bound", "l", 0, "Lower Bound Of Allowed Port Range")
	res.Command.PersistentFlags().IntVarP(&res.Config.AllowedPortUpperBound, "allowed-port-upper-bound", "u", 65535, "Upper Bound Of Allowed Port Range")
	res.Command.PersistentFlags().StringSliceVar(&res.Config.Tags, "tag", []string{}, "Tags")
	res.Command.PersistentFlags().StringVarP(&res.Config.FileBrowserDir, "dir", "f", "/", "Filebrowser Default Directory, Default=/")

	res.Command.Run = func(cmd *cobra.Command, args []string) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c := client.NewClient(res.Config.ServerIp, res.Config.ServerPort, res.Config.AllowedPortLowerBound, res.Config.AllowedPortUpperBound, res.Config.Tags, nil)
		c.FilebrowserDefaultDir = c.FilebrowserDefaultDir
		c.Start()
		wg.Wait()
	}

	return &res
}
