package main

import (
	"fmt"
	"time"

	"github.com/tigertony2536/go-line-notify/notiApp/notification"
)

func main() {
	// tasks, _ := notification.GetThisWeekTasks()
	// fmt.Println(notification.Format(tasks))
	// err := notification.NotifyWeekTasks()
	// if err == nil {
	// 	fmt.Println("Something wrong during notify weekly task0", err)
	// }
	for {
		fmt.Println("Goline is running...")
		if notification.ValidateWeek() {
			err := notification.NotifyWeekTasks()
			if err != nil {
				fmt.Println("Something wrong during notify weekly task0", err)
			}
		}
		if notification.ValidateDay() {
			err := notification.NotifyTodayTasks()
			if err != nil {
				fmt.Println("Something wrong during notify daily task", err)
			}
		}
		time.Sleep(1 * time.Second)
	}

}
