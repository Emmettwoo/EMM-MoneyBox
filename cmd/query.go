package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var queryCmd = &cobra.Command{
	Use:   "query {type} [condition]",
	Short: "Query for CashFlow data.",
	Long: `
Query for CashFlow data.

Types:
  date  ->  query by date, pass condition with format 19700101.
  id    ->  query by id, pass condition like 6241a836c72c65d9f343d891.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) <= 0 {
			return errors.New("must give a query type")
		}

		var queryType = args[0]
		switch queryType {
		case "id":

			if len(args) < 2 {
				return errors.New("must give a id string")
			}

			objectId, err := primitive.ObjectIDFromHex(args[1])
			if err != nil {
				panic(err)
			}

			cashFlow := model.GetCashFlowByObjectId(objectId)
			fmt.Println("cashFlow ", 0, ": ", cashFlow)

		case "date":
			var queryDate = time.Now()
			if len(args) > 1 {
				queryDate = util.FormatDateFromString(args[1])
			}

			// date format is yyyymmdd
			dayFlow := model.GetDayFlowByDate(queryDate)
			if dayFlow.IsEmpty() {
				fmt.Println("The day's flow is empty.")
			} else {
				cashFlowArray := model.GetCashFlowsByObjectIdArray(dayFlow.CashFlows)
				for index, cashFlow := range cashFlowArray {
					fmt.Println("cashFlow ", index, ": ", cashFlow)
				}
			}

		default:
			fmt.Println("Not supported query type.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
