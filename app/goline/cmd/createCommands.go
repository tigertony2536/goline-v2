/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/goline-v2/repository"
)

// insertCommandsCmd represents the insertCommands command
var insertCommandsCmd = &cobra.Command{
	Use:   "create",
	Short: "Insert a task to database",
	Long: `Insert a task to database
			InsertTask [TaskName] [Date] [Time]
			Date: YYYY-MM-DD
			Time: HH:MM:SS`,
	Run: func(cmd *cobra.Command, args []string) {
		task := repository.Task{Name: args[0], Date: args[1], Time: args[2]}
		id, err := repository.CreateTask(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Insert tasks successfully. Task's ID is %v", id)
	},
}

func init() {
	rootCmd.AddCommand(insertCommandsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// insertCommandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// insertCommandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
