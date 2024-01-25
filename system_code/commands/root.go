/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "letgo",
	Short: "letgo 是一个方便的本地go语言刷题工具",
	Long:  `描述： letgo 是一个 力扣/leetcode 本地刷题工具[目前仅支持go语言]。代码主要由 go 语言实现，让用户可以方便的选择自己喜欢的 IDE 进行刷题。`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd_middleware *cobra.Command, args []string) { },
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			panic(err.Error())
		}
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.system_code.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
