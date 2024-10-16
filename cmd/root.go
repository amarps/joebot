package cmd

import (
	"github.com/amarps/joebot/cmd/client"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "joebot",
		Short: "Command & Control Server/Client For Managing Machines Via Web Interface",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	clientCommand := client.Init()
	rootCmd.AddCommand(clientCommand.Command)
	rootCmd.AddCommand(initCmd)
}
