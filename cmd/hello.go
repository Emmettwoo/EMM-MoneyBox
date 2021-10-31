package cmd

import (
	"fmt"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "print Hello World!",
	Long:  `Say Hello to the World!`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Hello Wolrd!")
		// TODO: Test
		newCashFlowId := model.InsertCashFlowByEntity(model.CashFlowEntity{
			Amount:   1.23,
			Category: "test",
			Desc:     "this is a test.",
			Remark:   "null",
		})

		fmt.Println(model.GetCashFlowByObjectId(newCashFlowId))
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
