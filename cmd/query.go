package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var queryCmd = &cobra.Command{
	Use:   "query {type} [condition]",
	Short: "query for cash_flow data",
	Long: `
query for cash_flow data.

Types:
  id           ->  query by id, pass condition like 6241a836c72c65d9f343d891.
  date         ->  query by date, pass condition with format 19700101.
  desc_exact   ->  query by cashFlow describe information exact match.
  desc_fuzzy   ->  query by cashFlow describe information fuzzy match.`,
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

			cashFlow := cashFlowMapper.GetCashFlowByObjectId(objectId)
			fmt.Println("cashFlow ", 0, ": ", cashFlow.ToString())

		case "date":
			var queryDate = time.Now()
			if len(args) > 1 {
				queryDate = util.FormatDateFromString(args[1])
			}

			cashFlowList := cashFlowMapper.GetCashFlowsByBelongsDate(queryDate)

			if len(cashFlowList) == 0 {
				fmt.Println("the day's flow is empty")
			} else {
				for index, cashFlow := range cashFlowList {
					fmt.Println("cash_flow ", index, ": ", cashFlow.ToString())
				}
			}

		case "desc":
			fmt.Println("please try with desc_exact or desc_fuzzy instead")

		case "desc_exact":
			if len(args) < 2 {
				return errors.New("must give a desc string")
			}

			matchedCashFlow := cashFlowMapper.GetCashFlowsByExactDesc(args[1])
			if len(matchedCashFlow) == 0 {
				fmt.Println("no matched cash_flows")
				return nil
			}

			for index, cashFlow := range matchedCashFlow {
				fmt.Println("cash_flow ", index, ": ", cashFlow.ToString())
			}

		case "desc_fuzzy":
			if len(args) < 2 {
				return errors.New("must give a desc string")
			}

			matchedCashFlow := cashFlowMapper.GetCashFlowsByFuzzyDesc(args[1])
			if len(matchedCashFlow) == 0 {
				fmt.Println("no matched cash_flows")
				return nil
			}

			for index, cashFlow := range matchedCashFlow {
				fmt.Println("cash_flow ", index, ": ", cashFlow.ToString())
			}

		default:
			fmt.Println("Not supported query type")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}
