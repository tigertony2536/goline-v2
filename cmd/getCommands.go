/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/go-line-notify/controller"
	"github.com/tigertony2536/go-line-notify/model"
)

// getTasksCmd represents the getTasks command
var getTasksByDateCmd = &cobra.Command{
	Use:   "getTasksByDate",
	Args:  cobra.ExactArgs(2),
	Short: "Get tasks From Date-to-Date",
	Long: `Get Task from a specific period
	Parameter: Start(YYYY-MM-DD),End(YYYY-MM-DD) string `,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := model.GetByDate(args[0], args[1])
		if err != nil {
			log.Fatalln("Something wrong with query data")
		}
		s := controller.Format(tasks)
		fmt.Print(s)
	},
}

func init() {
	rootCmd.AddCommand(getTasksByDateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
