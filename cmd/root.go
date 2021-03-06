package cmd

import (
	"fmt"

	"github.com/emmettwoo/EMM-MoneyBox/mapper"
	"github.com/spf13/cobra"
)

var dayFlowMapper mapper.DayFlowMapper
var cashFlowMapper mapper.CashFlowMapper

var rootCmd = &cobra.Command{
	Use:   "EMM-MoneyBox",
	Short: "Root Commond",
	Long:  `Welcome to EMM-MoneyBox.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to EMM-MoneyBox.")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cashFlowMapper = mapper.GetCashFlowMapper()
	dayFlowMapper = mapper.GetDayFlowMapper()

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
