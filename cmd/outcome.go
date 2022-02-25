package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
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

		amount, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return err
		}
		// 取小數點後兩位
		amount, _ = decimal.NewFromFloat(amount).Round(2).Float64()

		category := args[1]
		desc := "null"
		if len(args) >= 3 {
			desc = args[2]
		}

		date := time.Time{}
		if len(args) >= 4 {
			var dateLayoutFormat = "20060102"
			date, _ = time.Parse(dateLayoutFormat, args[3])
		}

		newCashFlowId := model.InsertCashFlowByEntity(model.CashFlowEntity{
			Amount:   amount,
			Category: category,
			Desc:     desc,
			Remark:   "null",
		}, date)

		fmt.Println(model.GetCashFlowByObjectId(newCashFlowId))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(outcomeCmd)
}
