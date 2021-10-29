package cmd

import (
	"fmt"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "EMM-MoneyBox",
	Short: "Root Commond",
	Long:  `Welcome to EMM-MoneyBox.`,
	// 不带参数访问，默认打印今日dayFLow及其cashFlow
	Run: func(cmd *cobra.Command, args []string) {
		// 尝试获取今日dayFlow
		todayFlow := model.GetDayFlowByDate(time.Now())
		fmt.Println("todayFlow: ", todayFlow)

		// 用dayFlow换取cashFlow
		cashFlowArray := model.GetCashFlowsByObjectIdArray(todayFlow.CashFlows)
		for index, cashFlow := range cashFlowArray {
			fmt.Println("cashFlow ", index, ": ", cashFlow)
		}
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
