package server_cmd

import (
	"github.com/emmettwoo/EMM-MoneyBox/controller"
	"github.com/spf13/cobra"
)

var port int32

var startCmd4ApiServer = &cobra.Command{
	Use:   "start",
	Short: "start the api server",
	Run: func(cmd *cobra.Command, args []string) {
		controller.StartServer(port)
	},
}

func init() {
	startCmd4ApiServer.Flags().Int32VarP(
		&port, "port", "p", 8080, "api server port, default 8080")
	ServerCmd.AddCommand(startCmd4ApiServer)
}
