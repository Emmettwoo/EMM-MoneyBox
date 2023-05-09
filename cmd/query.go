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
	Short: "Query for CashFlow data.",
	Long: `
Query for CashFlow data.

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

			dayFlow := dayFlowMapper.GetDayFlowByDate(queryDate)

			if dayFlow.IsEmpty() {
				fmt.Println("The day's flow is empty.")
			} else {
				flowRefArray := flowRefMapper.GetFlowRefByDayFlowId(dayFlow.Id)
				for index, flowRef := range flowRefArray {
					fmt.Println("cashFlow ", index, ": ", cashFlowMapper.GetCashFlowByObjectId(flowRef.CashFlowId).ToString())
				}
			}

		case "desc":
			fmt.Println("Please try with desc_exact or desc_fuzzy instead.")

		case "desc_exact":
			if len(args) < 2 {
				return errors.New("must give a desc string")
			}

			matchedCashFlow := cashFlowMapper.GetCashFlowsByExactDesc(args[1])
			if len(matchedCashFlow) == 0 {
				fmt.Println("No Matched CashFlows.")
				return nil
			}

			for index, cashFlow := range matchedCashFlow {
				fmt.Println("cashFlow ", index, ": ", cashFlow.ToString())
			}

		case "desc_fuzzy":
			if len(args) < 2 {
				return errors.New("must give a desc string")
			}

			matchedCashFlow := cashFlowMapper.GetCashFlowsByFuzzyDesc(args[1])
			if len(matchedCashFlow) == 0 {
				fmt.Println("No Matched CashFlows.")
				return nil
			}

			for index, cashFlow := range matchedCashFlow {
				fmt.Println("cashFlow ", index, ": ", cashFlow.ToString())
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
