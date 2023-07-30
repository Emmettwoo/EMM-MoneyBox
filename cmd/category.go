package cmd

import (
	"errors"
	"github.com/emmettwoo/EMM-MoneyBox/service/category_service"
	"github.com/spf13/cobra"
)

var plainId4Category string
var parentPlainId4Category string
var categoryName4Category string

var queryCmd4Category = &cobra.Command{
	Use:   "query",
	Short: "query for category data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return category_service.QueryService(plainId4Category, categoryName4Category)
	},
}

var createCmd4Category = &cobra.Command{
	Use:   "create",
	Short: "create new category",
	RunE: func(cmd *cobra.Command, args []string) error {
		return category_service.CreateService(parentPlainId4Category, categoryName4Category)
	},
}

var deleteCmd4Category = &cobra.Command{
	Use:   "delete",
	Short: "delete category data",
	RunE: func(cmd *cobra.Command, args []string) error {
		return category_service.DeleteService(plainId4Category, categoryName4Category)
	},
}

var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "operating category data",
	Long: `
Operating category data by several sub-commands.
Provide sub-commands: [query, create, delete].`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}

func init() {

	// add sub-command: query
	queryCmd4Category.Flags().StringVarP(
		&plainId4Category, "id", "i", "", "query by id")
	queryCmd4Category.Flags().StringVarP(
		&categoryName4Category, "name", "n", "", "query by name")
	categoryCmd.AddCommand(queryCmd4Category)

	// add sub-command: create
	createCmd4Category.Flags().StringVarP(
		&parentPlainId4Category, "parent", "p", "", "category's parent's id (optional)")
	createCmd4Category.Flags().StringVarP(
		&categoryName4Category, "name", "n", "", "category's name (required)")
	categoryCmd.AddCommand(createCmd4Category)

	// add sub-command: delete
	deleteCmd4Category.Flags().StringVarP(
		&plainId4Category, "id", "i", "", "delete by id")
	deleteCmd4Category.Flags().StringVarP(
		&categoryName4Category, "name", "n", "", "delete by name")
	categoryCmd.AddCommand(deleteCmd4Category)

	// todo: add sub-command: update (by id)

	// add command: category
	rootCmd.AddCommand(categoryCmd)
}
