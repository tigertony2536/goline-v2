package controller

import (
	"fmt"
	"log"
)

func SendNotiApp() {
	// server := controller.NewServer(":8080")
	// log.Fatal(server.Start())

	weekNoti, err := GetWeeklyNoti()
	if err != nil {
		log.Fatal(err)
	}
	dailyNoti, err := GetDailyNoti()
	if err != nil {
		log.Fatal(err)
	}
	if weekNoti.Tasks != nil {
		resp, err := SendNotification(weekNoti)
		fmt.Print(resp)
		if err != nil {
			log.Fatal(err)
		}
	}

	if dailyNoti.Tasks != nil {
		resp, err := SendNotification(weekNoti)
		fmt.Print(resp)
		if err != nil {
			log.Fatal(err)
		}
	}
}
