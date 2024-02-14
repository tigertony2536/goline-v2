/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/abdfnx/gosh"
	"github.com/spf13/cobra"
	"github.com/tigertony2536/goline-v2/app/notiApp/notification"
)

// notifyCommandsCmd represents the notifyCommands command
var notifyCommandsCmd = &cobra.Command{
	Use:   "notify",
	Short: "Send notify to line api",
	Long: `Send daily and weekly tasks to line chat if exist 
	(use this command with task schelduler)`,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	err := notification.NotifyTodayTasks()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	err = notification.NotifyWeekTasks()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println("Send notify successfully")
	// },
}

var startCommandsCmd = &cobra.Command{
	Use:   "start",
	Short: "Start notify app",
	Long:  `Start notify app`,
	Run: func(cmd *cobra.Command, args []string) {
		command1 := exec.Command("notiApp")
		err := command1.Start()
		if err != nil {
			log.Fatal("something wrong during start notify app:", err)
		}
		fmt.Println("Start notify successfully")
	},
}

var stopCommandsCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop notify app",
	Long:  `Stop notify app`,
	Run: func(cmd *cobra.Command, args []string) {
		err, _, _ := gosh.PowershellOutput("taskkill /IM notiApp.exe /F")
		if err != nil {
			log.Fatal("something wrong during Stop notify app:", err)
		}
		fmt.Println("Stop notify successfully")
	},
}

var statusCommandsCmd = &cobra.Command{
	Use:   "status",
	Short: "Check notification status",
	Long:  `Check notification status`,
	Run: func(cmd *cobra.Command, args []string) {
		err, _, _ := gosh.PowershellOutput("Get-Process notiApp")
		// fmt.Println(a)
		// fmt.Println(b)
		if err != nil {
			fmt.Println("Notification is off")
		} else {
			fmt.Println("Notification is on")
		}
	},
}

var nowCommandsCmd = &cobra.Command{
	Use:   "now",
	Short: "Send recent day or week task notification immediately",
	Long:  `Send recent day or week task notification immediately`,
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
	notifyCommandsCmd.AddCommand(startCommandsCmd)
	notifyCommandsCmd.AddCommand(stopCommandsCmd)
	notifyCommandsCmd.AddCommand(statusCommandsCmd)
	notifyCommandsCmd.AddCommand(nowCommandsCmd)

}
