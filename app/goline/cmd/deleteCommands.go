/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	model "github.com/tigertony2536/goline-v2/repository"
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

}
