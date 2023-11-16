package cmd

import (
	"errors"
	"github.com/emmettwoo/EMM-MoneyBox/service/cash_flow_service"
	"github.com/spf13/cobra"
)

var amount4CashFlow float64
var belongsDate4CashFlow string
var categoryName4CashFlow string
var exactDescription4CashFlow string
var fuzzyDescription4CashFlow string
var plainId4CashFlow string

var queryCmd4CashFlow = &cobra.Command{
	Use:   "query",
	Short: "query for cash_flow data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cash_flow_service.QueryService(
			plainId4CashFlow, belongsDate4CashFlow, exactDescription4CashFlow, fuzzyDescription4CashFlow)
	},
}

var deleteCmd4CashFlow = &cobra.Command{
	Use:   "delete",
	Short: "delete cash_flow by specific type",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cash_flow_service.DeleteService(plainId4CashFlow, belongsDate4CashFlow)
	},
}

var outcomeCmd4CashFlow = &cobra.Command{
	Use:   "outcome",
	Short: "add new outcome cash_flow",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cash_flow_service.OutcomeService(
			belongsDate4CashFlow, categoryName4CashFlow, amount4CashFlow, exactDescription4CashFlow)
	},
}

var cashCmd = &cobra.Command{
	Use:   "cash",
	Short: "operating cash_flow data",
	Long: `
Operating cash data by several sub-commands.
Provide sub-commands: [query, delete, outcome].`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}

func init() {

	// add sub-command: query
	queryCmd4CashFlow.Flags().StringVarP(
		&plainId4CashFlow, "id", "i", "", "query by id")
	queryCmd4CashFlow.Flags().StringVarP(
		&belongsDate4CashFlow, "date", "b", "", "query by belongs-date")
	queryCmd4CashFlow.Flags().StringVarP(
		&exactDescription4CashFlow, "exact", "e", "", "query by exact-description")
	queryCmd4CashFlow.Flags().StringVarP(
		&fuzzyDescription4CashFlow, "fuzzy", "f", "", "query by fuzzy-description")
	cashCmd.AddCommand(queryCmd4CashFlow)

	// add sub-command: delete
	deleteCmd4CashFlow.Flags().StringVarP(
		&plainId4CashFlow, "id", "i", "", "delete by id")
	deleteCmd4CashFlow.Flags().StringVarP(
		&belongsDate4CashFlow, "date", "b", "", "delete by belongs-date")
	cashCmd.AddCommand(deleteCmd4CashFlow)

	// todo: add sub-command: update (by id)

	// add sub-command: outcome
	outcomeCmd4CashFlow.Flags().StringVarP(
		&belongsDate4CashFlow, "date", "b", "", "flow's belongs-date (optional, blank for today)")
	outcomeCmd4CashFlow.Flags().StringVarP(
		&categoryName4CashFlow, "category", "c", "", "flow's category name (required)")
	outcomeCmd4CashFlow.Flags().Float64VarP(
		&amount4CashFlow, "amount", "a", 0.00, "flow's amount (required)")
	outcomeCmd4CashFlow.Flags().StringVarP(
		&exactDescription4CashFlow, "description", "d", "", "flow's description (optional, could be blank)")

	cashCmd.AddCommand(outcomeCmd4CashFlow)

	// add command: cash
	rootCmd.AddCommand(cashCmd)
}
