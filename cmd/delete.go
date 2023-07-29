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
	Short: "delete cash_flow by specific type",
	Long: `
Delete cash_flow in specific type.
Types:
  id  ->  delete by objectId like 621b18e3cbab6c2f4d75d0cb, could pass multi value.
  date -> delete all cash_flow which belongs to the date like 19700101.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		if len(args) <= 0 {
			return errors.New("must give a delete type")
		}

		var queryType = args[0]
		switch queryType {

		case "id":
			// check number of record to be deleted.
			totalToBeDelete := len(args) - 1
			if totalToBeDelete <= 0 {
				return errors.New("must give at lease one cash_flow id")
			}

			currentDelete := 1
			for currentDelete <= totalToBeDelete {
				objectId, err := primitive.ObjectIDFromHex(args[currentDelete])
				if err != nil {
					util.Logger.Errorw("hex to object id failed",
						"objectId", args[currentDelete], "error", err)
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

			cashFlowArray := cashFlowMapper.DeleteCashFlowByBelongsDate(deleteDate)
			for index, cashFlow := range cashFlowArray {
				fmt.Println("cash_flow ", index, ": ", cashFlow)
			}

		default:
			fmt.Println("not supported delete type")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
