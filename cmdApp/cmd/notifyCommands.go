/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/go-line-notify/notiApp/notification"
)

// notifyCommandsCmd represents the notifyCommands command
var notifyCommandsCmd = &cobra.Command{
	Use:   "notify",
	Short: "Send notify to line api",
	Long: `Send daily and weekly tasks to line chat if exist 
	(use this command with task schelduler)`,
	Run: func(cmd *cobra.Command, args []string) {
		err := notification.NotifyTodayTasks()
		if err != nil {
			log.Fatal(err)
		}
		err = notification.NotifyWeekTasks()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Send notify successfully")
	},
}

func init() {
	rootCmd.AddCommand(notifyCommandsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// notifyCommandsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// notifyCommandsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
