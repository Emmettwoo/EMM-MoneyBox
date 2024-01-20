package cmd

import (
	"fmt"

	"github.com/emmettwoo/EMM-MoneyBox/cmd/cash_flow_cmd"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "EMM-MoneyBox",
	Short: "root command",
	Long:  `Welcome to EMM-MoneyBox.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to EMM-MoneyBox.")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// add command: cash
	rootCmd.AddCommand(cash_flow_cmd.CashCmd)
}
