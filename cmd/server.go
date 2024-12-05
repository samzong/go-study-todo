// Package cmd defines the commands for the CLI application
package cmd

import (
	"strings"
	api "todo/api"

	"github.com/spf13/cobra"
)

// serverCmd represents the 'server' command in the CLI application
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the API server",
	Long:  `Start, the API Server to manage tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		// Start the API server
		port := strings.TrimSpace(cmd.Flag("port").Value.String())
		address := strings.TrimSpace(cmd.Flag("address").Value.String())
		api.StartServer(address, port)
	},
}

func init() {
	// Add the 'server' command
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("address", "", "localhost", "Server address")
	serverCmd.Flags().StringP("port", "", "8080", "Server port")
}
