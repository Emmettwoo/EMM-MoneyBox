package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/emmettwoo/EMM-MoneyBox/model"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "EMM-MoneyBox",
	Short: "Basic Commond",
	Long:  `Welcome to EMM-MoneyBox.`,
	Run: func(cmd *cobra.Command, args []string) {
		// 尝试获取今日dayFlow
		todayFlow := model.GetDayFlowByDate(time.Now())
		fmt.Println("todayFlow: ", todayFlow)

		// 用dayFlow换取cashFlow
		cashFlowArray := model.GetCashFlowsByObjectIdArray(todayFlow.CashFlows)
		for index, cashFlow := range cashFlowArray {
			fmt.Println("cashFlow ", index, ": ", cashFlow)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.EMM-MoneyBox.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".EMM-MoneyBox" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".EMM-MoneyBox")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// Print config file path.
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
