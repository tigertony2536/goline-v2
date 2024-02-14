/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/goline-v2/repository"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update task by id",
	Long:  `Update task by id`,
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
		err = repository.UpdateTask(id, nameFlag, dateFlag, timeFlag)
		if err != nil {
			log.Fatal("Something wrong durring updating task\n", err)
		}

		fmt.Println("updated task successfully")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	var name string
	updateCmd.Flags().StringVarP(&name, "name", "n", "", "update a task by name")
	var date string
	updateCmd.Flags().StringVarP(&date, "date", "d", "", "update a task by name")
	var time string
	updateCmd.Flags().StringVarP(&time, "time", "t", "", "update a task by name")
}
