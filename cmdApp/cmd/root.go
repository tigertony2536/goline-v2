/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goline",
	Short: "A Task Scheduler app implement with Go",
	Long: `goline is a task scheduler app. You can create, read, 
	update and delete tasks. Setup time for it to send you line notifications:`,
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	// for {
	// 	if len(os.Args) > 1 {
	// 		go fmt.Println(os.Args[1:])
	// 	}

	// }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
