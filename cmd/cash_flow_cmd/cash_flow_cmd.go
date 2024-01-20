package cash_flow_cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var plainId string
var amount float64
var belongsDate string
var categoryName string
var descriptionExact string
var descriptionFuzzy string

var CashCmd = &cobra.Command{
	Use:   "cash",
	Short: "operating cash_flow data",
	Long: `
Operating cash data by several sub-commands.
Provide sub-commands: [query, delete, outcome].`,

	// todo(emmett): add sub-command: update by id
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}
