package category_cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var plainId string
var parentPlainId string
var categoryName string

var CategoryCmd = &cobra.Command{
	Use:   "category",
	Short: "operating category data",
	Long: `
Operating category data by several sub-commands.
Provide sub-commands: [query, create, delete].`,

	// todo(emmett): add sub-command: update by id
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}
