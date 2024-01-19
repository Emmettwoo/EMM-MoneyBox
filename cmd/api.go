package cmd

import (
	"errors"

	"github.com/emmettwoo/EMM-MoneyBox/api"
	"github.com/spf13/cobra"
)

var port4ApiServer int32

var startCmd4ApiServer = &cobra.Command{
	Use:   "start",
	Short: "start the api server",
	Run: func(cmd *cobra.Command, args []string) {
		api.StartServer(port4ApiServer)
	},
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "manage api server",
	Long: `
Managing application API server.
Provide sub-commands: [start].`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}

func init() {
	// add sub-command: start
	startCmd4ApiServer.Flags().Int32VarP(
		&port4ApiServer, "port", "p", 8080, "api server port, default 8080")
	apiCmd.AddCommand(startCmd4ApiServer)

	// add command: cash
	rootCmd.AddCommand(apiCmd)
}
