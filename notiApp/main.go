package main

import (
	"fmt"
	"time"

	"github.com/tigertony2536/Goline/notiApp/notification"
)

func main() {
	fmt.Println("Goline is running...")
	for {
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
