package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "print Hello World!",
	Long:  `Say Hello to the World!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Wolrd!")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
