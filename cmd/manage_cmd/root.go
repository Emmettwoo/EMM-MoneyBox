package manage_cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var fromDate string
var toDate string
var filePath string

var ManageCmd = &cobra.Command{
	Use:   "manage",
	Short: "manage setting and data",
	Long: `
Managing program setting and data by several sub-commands.
Provide sub-commands: [export, import].`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}
