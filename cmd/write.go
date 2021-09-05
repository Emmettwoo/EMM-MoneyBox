package cmd

import (
	"fmt"

	"github.com/emmettwoo/EMM-MoneyBox/util"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "write yyyy-mm-dd title amount [desc]",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 3 {
			util.Write(args[0], args[1], util.String2Float32(args[2]), "")
		} else if len(args) == 4 {
			util.Write(args[0], args[1], util.String2Float32(args[2]), args[3])
		} else {
			fmt.Println("The number of parameters does not match!")
		}
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)
}
