/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/goline/model"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task by id",
	Long:  `Update task by id:`,
	Run: func(cmd *cobra.Command, args []string) {
		nameFlag, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal("Something wrong durring fetch name's flag", err)
		}
		dateFlag, err := cmd.Flags().GetString("date")
		if err != nil {
			log.Fatal("Something wrong durring fetch date's flag", err)
		}
		timeFlag, err := cmd.Flags().GetString("time")
		if err != nil {
			log.Fatal("Something wrong durring fetch time's flag", err)
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal("Something wrong durring convert string to int", err)
		}
		err = model.UpdateTask(id, nameFlag, dateFlag, timeFlag)
		if err != nil {
			log.Fatal("Something wrong durring updating task\n", err)
		}
		fmt.Println("updated task successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().String("name", "", "use name as operation criteria")
	updateCmd.Flags().String("date", "", "use date as operation criteria")
	updateCmd.Flags().String("time", "", "use time as operation criteria")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
