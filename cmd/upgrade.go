package cmd

import (
	"fmt"
	"strconv"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade data version.",
	Long:  ` for test only. `,
	RunE: func(cmd *cobra.Command, args []string) error {

		var queryDateNumber = 20221120

		for queryDateNumber < 20221130{
			var queryDate = util.FormatDateFromString(strconv.Itoa(queryDateNumber))
			dayFlow := dayFlowMapper.GetDayFlowByDate(queryDate)
			if dayFlow.IsEmpty() {
				fmt.Printf("%d's flow is empty.\n", queryDateNumber)
			} else {
				fmt.Printf("%d's flow as below.\n", queryDateNumber)
				cashFlowArray := cashFlowMapper.GetCashFlowsByObjectIdArray(dayFlow.CashFlows)
				for index, cashFlow := range cashFlowArray {
					fmt.Println("cashFlow", index, ": ", cashFlow.ToString())

					// todo: 轉換前先調整查詢和插入的邏輯，與關聯信息建立聯係。
					// 把目前已有的數據轉換出關聯信息
					_ = flowRefMapper.InsertFlowRefByEntity(entity.FlowRefEntity{
						DayFlowId:   dayFlow.Id,
						CashFlowId: cashFlow.Id,
					})
				}
				// 刪除原 day_flow 裏的關聯關係
				dayFlow.CashFlows = nil
				dayFlowMapper.UpdateDayFlowByEntity(dayFlow)
			}
			queryDateNumber++
		}

		return nil
	},
}

func init() {
	// 臨時升級數據結構的指令，升級完成后會刪除
	rootCmd.AddCommand(upgradeCmd)
}
