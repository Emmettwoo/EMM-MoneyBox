package cmd

import (
	"errors"
	"fmt"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var deleteCmd = &cobra.Command{
	Use:   "delete {type} {...params}",
	Short: "delete CashFlow in specific type.",
	Long: `
Delete CashFlow in specific type.
Types:
  id  ->  delete by objectId like 621b18e3cbab6c2f4d75d0cb.
  date -> delete all CashFlow which belongs to the date like 19700101.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) <= 0 {
			return errors.New("must give a delete type")
		}

		var queryType = args[0]
		switch queryType {

		case "id":

			totalToBeDelete := len(args) - 1
			if totalToBeDelete <= 0 {
				return errors.New("must give at lease one delete id")
			}

			currentDelete := 1
			for currentDelete <= totalToBeDelete {
				objectId, err := primitive.ObjectIDFromHex(args[currentDelete])
				if err != nil {
					panic(err)
				}

				cashFlow := cashFlowMapper.DeleteCashFlowByObjectId(objectId)
				fmt.Println("cashFlow ", currentDelete-1, ": ", cashFlow)
				currentDelete++
			}

		case "date":

			var deleteDate = time.Now()
			if len(args) > 1 {
				deleteDate = util.FormatDateFromString(args[1])
			}

			// date format is yyyymmdd
			dayFlow := dayFlowMapper.DeleteDayFlowByDate(deleteDate)
			if dayFlow.IsEmpty() || len(dayFlow.CashFlows) == 0 {
				fmt.Println("The day's flow is empty.")
			} else {
				cashFlowArray := cashFlowMapper.GetCashFlowsByObjectIdArray(dayFlow.CashFlows)
				for index, cashFlow := range cashFlowArray {
					fmt.Println("cashFlow ", index, ": ", cashFlow)
				}
			}

		default:
			fmt.Println("Not supported delete type.")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
