/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/goline/model"
	"github.com/tigertony2536/goline/notiApp/notification"
)

// getTasksCmd represents the getTasks command
var getTaskCmd = &cobra.Command{
	Use:   "getTask",
	Short: "Get existing tasks in app",
	Long: `Get existing tasks in app. User have to specify criteria with subcommand
	getTask date <YYYY-MM-DD> <YYYY-MM-DD>
	getTask id <id>
	getTask name <pattern>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please specify subcommandd and arguments.")
			fmt.Println(" -> date <YYYY-MM-DD> <YYYY-MM-DD>")
			fmt.Println(" -> id <id>")
			fmt.Println(" -> name <pattern>")
		}
	},
}

var dateCmd = &cobra.Command{
	Use:   "date",
	Args:  cobra.ExactArgs(2),
	Short: "Get task by date",
	Long: `Get tasks scheduled between a specific period
	getTask date <YYYY-MM-DD> <YYYY-MM-DD>`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := model.GetByDate(args[0], args[1])
		if err != nil {
			log.Fatalln("Something wrong with query data")
		}
		s := notification.Format(tasks)
		fmt.Print(s)
	},
}

var idCmd = &cobra.Command{
	Use:   "id",
	Args:  cobra.ExactArgs(1),
	Short: "Get task by id",
	Long: `Get tasks by task's id
	getTask id <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatal(err)
		}
		tasks, err := model.GetByID(id)
		if err != nil {
			log.Fatalln("Something wrong with query data")
		}
		s := notification.Format(tasks)
		fmt.Print(s)
	},
}

var nameCmd = &cobra.Command{
	Use:   "name",
	Args:  cobra.ExactArgs(1),
	Short: "Get task by name",
	Long: `Get tasks by task's name
	getTask name <pattern>`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := model.GetByName(args[0])
		if err != nil {
			log.Fatalln("Something wrong with query data")
		}
		s := notification.Format(tasks)
		fmt.Print(s)
	},
}

var allCmd = &cobra.Command{
	Use:   "all",
	Args:  cobra.NoArgs,
	Short: "Get all tasks",
	Long:  `Get all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := model.GetAllTasks()
		if err != nil {
			log.Fatalln("Something wrong with query data: ", err)
		}
		s := notification.Format(tasks)
		fmt.Print(s)
	},
}

func init() {
	rootCmd.AddCommand(getTaskCmd)

	getTaskCmd.AddCommand(dateCmd)
	getTaskCmd.AddCommand(idCmd)
	getTaskCmd.AddCommand(nameCmd)
	getTaskCmd.AddCommand(allCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getTasksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getTasksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
