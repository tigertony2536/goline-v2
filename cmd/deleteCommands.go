/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/go-line-notify/model"
)

// deleteCommandsCmd represents the deleteCommands command
var deleteCommandsCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a Task by id",
	Long:  `Delete a Task by id. You can get task id by get all task command`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, id := range args {
			n, err := strconv.Atoi(id)
			if err != nil {
				log.Fatal(err)
			}
			err = model.DeleteTask(n)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("Delete task id %v successfully\n", id)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCommandsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCommandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCommandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
