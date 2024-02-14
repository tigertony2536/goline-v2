/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/tigertony2536/goline-v2/app/notiApp/notification"
	"github.com/tigertony2536/goline-v2/repository"
)

// getTasksCmd represents the getTasks command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get existing tasks in app",
	Long: `Get existing tasks in app. User have to specify criteria with subcommand
	getTask date <YYYY-MM-DD> <YYYY-MM-DD>
	getTask id <id>
	getTask name <pattern>`,
	Run: func(cmd *cobra.Command, args []string) {
		idFlag, _ := cmd.Flags().GetString("id")
		dateFlag, _ := cmd.Flags().GetString("date")
		nameFlag, _ := cmd.Flags().GetString("name")
		allFlag, _ := cmd.Flags().GetString("all")

		// Case ID Flag
		if idFlag != "" {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			tasks, err := repository.GetByID(id)
			if err != nil {
				log.Fatalln("Something wrong with query data")
			}
			s := notification.Format(tasks, "")
			fmt.Print(s)
		}

		// Case Name Flag
		if nameFlag != "" {
			tasks, err := repository.GetByName(args[0])
			if err != nil {
				log.Fatalln("Something wrong with query data")
			}
			s := notification.Format(tasks, "")
			fmt.Print(s)
		}
		// Case Date Flag
		if dateFlag != "" {
			tasks, err := repository.GetByDate(args[0], args[1])
			if err != nil {
				log.Fatalln("Something wrong with query data")
			}
			s := notification.Format(tasks, "")
			fmt.Print(s)
		}
		// Case All Flag
		if allFlag != "" {
			tasks, err := repository.GetAllTasks()
			if err != nil {
				log.Fatalln("Something wrong with query data: ", err)
			}
			s := notification.Format(tasks, "")
			fmt.Print(s)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	var id string
	getCmd.Flags().StringVarP(&id, "id", "i", "", "return a task by id")
	var name string
	getCmd.Flags().StringVarP(&name, "name", "n", "", "return a task by id")
	var date string
	getCmd.Flags().StringVarP(&date, "date", "d", "", "return a task by id")
	var all string
	getCmd.Flags().StringVarP(&all, "all", "a", "", "return a task by id")
	getCmd.Flags().Lookup("all").NoOptDefVal = "a"

}
