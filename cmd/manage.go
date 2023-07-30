package cmd

import (
	"errors"
	"github.com/emmettwoo/EMM-MoneyBox/service/manage_service"
	"github.com/spf13/cobra"
)

var fromDate string
var toDate string
var filePath string

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export data to excel",
	RunE: func(cmd *cobra.Command, args []string) error {
		return manage_service.ExportService(fromDate, toDate, filePath)
	},
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "import data from excel",
	RunE: func(cmd *cobra.Command, args []string) error {
		return manage_service.ImportService(filePath)
	},
}

var manageCmd = &cobra.Command{
	Use:   "manage",
	Short: "manage setting and data",
	Long: `
Managing program setting and data by several sub-commands.
Provide sub-commands: [export, import].`,

	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("must provide a valid sub command")
	},
}

func init() {

	// add sub-command: export
	exportCmd.Flags().StringVarP(&fromDate, "from", "f", "", "from date(include), e.x. 19700101")
	exportCmd.Flags().StringVarP(&toDate, "to", "t", "", "to date(include), e.x. 19700101")
	exportCmd.Flags().StringVarP(&filePath, "output", "o", "", "output path, default ./export.xlsx")
	manageCmd.AddCommand(exportCmd)

	// add sub-command: import
	importCmd.Flags().StringVarP(&filePath, "input", "i", "", "input path, e.x. ~/export.xlsx")
	manageCmd.AddCommand(importCmd)

	// add command: manage
	rootCmd.AddCommand(manageCmd)
}
