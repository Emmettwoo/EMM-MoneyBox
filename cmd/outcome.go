package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/entity"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
)

var outcomeCmd = &cobra.Command{
	Use:   "outcome {amount} {category} [date] [desc]",
	Short: "add new outcome cash_flow",
	Long: `
Insert one new outcome CashFlow.

Params:
  amount  ->  float value with most two digits after the dot.
  category  ->  category name that already created.
  date  ->  define cash_flow belongs date, current date by default.
  desc  ->  other describe string for this cash_flow.`,
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
		category := args[1]
		categoryEntity := CategoryMapper.GetCategoryByName(category)
		if categoryEntity.IsEmpty() {
			panic("category does not exist")
		}

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

		newCashFlowId := cashFlowMapper.InsertCashFlowByEntity(entity.CashFlowEntity{
			CategoryId:  categoryEntity.Id,
			BelongsDate: date,
			Type:        "OUTCOME",
			Amount:      amount,
			Desc:        desc,
			Remark:      "",
		})

		fmt.Println(cashFlowMapper.GetCashFlowByObjectId(newCashFlowId))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(outcomeCmd)
}
