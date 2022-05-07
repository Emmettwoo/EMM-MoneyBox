package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/emmettwoo/EMM-MoneyBox/mapper/mongodb"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var outcomeCmd = &cobra.Command{
	Use:   "outcome {amount} {category} [date] [desc]",
	Short: "outcome one new CashFlow.",
	Long: `
Insert one new outcome CashFlow.

Params:
  amount  ->  float value with two digits after the dot.
  category  ->  any string, for category purpose in the comming version.
  date  ->  define CashFlow belongs date, current date by default.
  desc  ->  other describe string for this CashFlow.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) < 2 {
			return errors.New("must provide amount and category params")
		}

		// 必填參數1: 金額
		amount, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}
		// 取小數點後兩位
		amount, _ = decimal.NewFromFloat(amount).Round(2).Float64()

		// 必填參數2: 類別
		// todo(emmett): 可以維護一個 category 列表，僅允許設定好的類別變數
		category := args[1]

		// 選填參數3: 日期（默認當天）
		date := time.Time{}
		if len(args) >= 3 {
			var dateLayoutFormat = "20060102"
			date, _ = time.Parse(dateLayoutFormat, args[2])
		}

		// 選填參數4: 描述（默認爲空）
		desc := ""
		if len(args) >= 4 {
			desc = args[3]
		}

		newCashFlowId := mongodb.InsertCashFlowByEntity(entity.CashFlowEntity{
			Amount:   amount,
			Category: category,
			Desc:     desc,
			Remark:   "",
		}, date)

		fmt.Println(mongodb.GetCashFlowByObjectId(newCashFlowId))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(outcomeCmd)
}
