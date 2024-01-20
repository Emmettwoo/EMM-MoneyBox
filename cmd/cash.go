package cmd

import (
	"errors"
	"fmt"

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

		// Valid params through command.
		if cash_flow_service.IsQueryFieldsConflicted(plainId4CashFlow, belongsDate4CashFlow, exactDescription4CashFlow, fuzzyDescription4CashFlow) {
			return errors.New("should have one and only one query type")
		}

		// if id is not empty, use it for query.
		if plainId4CashFlow != "" {
			cashFlowEntity, err := cash_flow_service.QueryById(plainId4CashFlow)
			if err != nil {
				return err
			}
			fmt.Println("cash_flow ", 0, ": ", cashFlowEntity.ToString())
			return nil
		}

		// else if date is not empty, use it for query.
		if belongsDate4CashFlow != "" {
			cashFlowEntityList, err := cash_flow_service.QueryByDate(belongsDate4CashFlow)
			if err != nil {
				return err
			}
			if len(cashFlowEntityList) == 0 {
				fmt.Println("the day's flow is empty")
				return nil
			}
			for index, cashFlowEntity := range cashFlowEntityList {
				fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
			}
			return nil
		}

		// else if exact_desc is not empty, use it for query.
		if exactDescription4CashFlow != "" {
			cashFlowEntityList, err := cash_flow_service.QueryByExactDescription(exactDescription4CashFlow)
			if err != nil {
				return err
			}
			if len(cashFlowEntityList) == 0 {
				fmt.Println("no matched cash_flows")
				return nil
			}

			for index, cashFlowEntity := range cashFlowEntityList {
				fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
			}
		}

		// else if fuzzy_desc is not empty, use it for query.
		if fuzzyDescription4CashFlow != "" {
			cashFlowEntityList, err := cash_flow_service.QueryByFuzzyDescription(fuzzyDescription4CashFlow)
			if err != nil {
				return err
			}
			if len(cashFlowEntityList) == 0 {
				fmt.Println("no matched cash_flows")
				return nil
			}

			for index, cashFlowEntity := range cashFlowEntityList {
				fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
			}
		}

		return errors.New("not supported query type")
	},
}

var deleteCmd4CashFlow = &cobra.Command{
	Use:   "delete",
	Short: "delete cash_flow by specific type",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Valid params through command.
		if cash_flow_service.IsDeleteFieldsConflicted(plainId4CashFlow, belongsDate4CashFlow) {
			return errors.New("should have one and only one delete type")
		}

		if plainId4CashFlow != "" {

			cashFlowEntity, err := cash_flow_service.DeleteById(plainId4CashFlow)
			if err != nil {
				return err
			}
			fmt.Println("cash_flow ", 0, ": ", cashFlowEntity.ToString())
			return nil
		}

		if belongsDate4CashFlow != "" {
			cashFlowEntityList, err := cash_flow_service.DeleteByDate(belongsDate4CashFlow)
			if err != nil {
				return err
			}
			if len(cashFlowEntityList) == 0 {
				fmt.Println("the day's flow is empty")
				return nil
			}
			for index, cashFlowEntity := range cashFlowEntityList {
				fmt.Println("cash_flow ", index, ": ", cashFlowEntity.ToString())
			}
			return nil
		}

		return errors.New("not supported delete type")
	},
}

var outcomeCmd4CashFlow = &cobra.Command{
	Use:   "outcome",
	Short: "add new outcome cash_flow",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !cash_flow_service.IsOutcomeRequiredFiledSatisfied(categoryName4CashFlow, amount4CashFlow) {
			return errors.New("some required fields are empty")
		}
		cashFlowEntity, err :=  cash_flow_service.SaveOutcome(belongsDate4CashFlow, categoryName4CashFlow, amount4CashFlow, exactDescription4CashFlow)
		if err != nil {
			return err
		}
		fmt.Println("cash_flow ", 0, ": ", cashFlowEntity.ToString())
		return nil
	},
}

var incomeCmd4CashFlow = &cobra.Command{
	Use:   "income",
	Short: "add new income cash_flow",
	RunE: func(cmd *cobra.Command, args []string) error {

		if !cash_flow_service.IsIncomeRequiredFiledSatisfied(categoryName4CashFlow, amount4CashFlow) {
			return errors.New("some required fields are empty")
		}
		cashFlowEntity, err := cash_flow_service.SaveIncome(belongsDate4CashFlow, categoryName4CashFlow, amount4CashFlow, exactDescription4CashFlow)
		if err != nil {
			return err
		}
		fmt.Println("cash_flow ", 0, ": ", cashFlowEntity.ToString())
		return nil
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

// todo(emmett): register everything in a route.go file instead.
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

	// add sub-command: income
	incomeCmd4CashFlow.Flags().StringVarP(
		&belongsDate4CashFlow, "date", "b", "", "flow's belongs-date (optional, blank for today)")
	incomeCmd4CashFlow.Flags().StringVarP(
		&categoryName4CashFlow, "category", "c", "", "flow's category name (required)")
	incomeCmd4CashFlow.Flags().Float64VarP(
		&amount4CashFlow, "amount", "a", 0.00, "flow's amount (required)")
	incomeCmd4CashFlow.Flags().StringVarP(
		&exactDescription4CashFlow, "description", "d", "", "flow's description (optional, could be blank)")
	cashCmd.AddCommand(incomeCmd4CashFlow)

	// add command: cash
	rootCmd.AddCommand(cashCmd)
}
