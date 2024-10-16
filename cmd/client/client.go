package client

import (
	"github.com/spf13/cobra"
)

type ClientCommand struct {
	Command        *cobra.Command
	ServerIp       string
	ServerPort     int
	FileBrowserDir string
}

func Init() *ClientCommand {
	res := ClientCommand{}
	res.Command = &cobra.Command{
		Use:   "client",
		Short: "Client Mode",
	}

	res.Command.PersistentFlags().StringVarP(&res.ServerIp, "license", "l", "", "name of license for the project")
	res.Command.PersistentFlags().IntVarP(&res.ServerPort, "license", "l", 0, "name of license for the project")
	res.Command.PersistentFlags().StringVarP(&res.FileBrowserDir, "license", "l", "", "name of license for the project")
	res.Command.PersistentFlags().Bool("viper", true, "use Viper for configuration")

	return &res
}
