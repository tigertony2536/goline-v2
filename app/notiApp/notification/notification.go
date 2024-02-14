package notification

import (
	"fmt"
	"strconv"
	"time"

	"github.com/tigertony2536/goline-v2/config"
	"github.com/tigertony2536/goline-v2/repository"
)

var (
	cfg config.AppConfig
)

func init() {
	cfg = config.GetAppConfig()
}

func getWeekDay() (time.Time, time.Time) {
	start := time.Now()

	for start.Weekday() != time.Monday {
		start = start.Add(-time.Hour * 24)
	}

	end := start.Add(time.Hour * 144)
	return start, end
}

func GetTodayTasks() (repository.TaskGroup, error) {
	today := time.Now().Format(time.DateOnly)

	noti, err := repository.GetByDate(today, today)
	if err != nil {
		return noti, err
	}

	return noti, nil
}

func GetThisWeekTasks() (repository.TaskGroup, error) {
	start, end := getWeekDay()

	noti, err := repository.GetByDate(start.Format(time.DateOnly), end.Format(time.DateOnly))
	if err != nil {
		return noti, err
	}

	return noti, nil
}

func Format(noti repository.TaskGroup, notiType string) string {
	//noti type 'daily' or 'weekly'
	s := "\n"
	if notiType != "" {
		s += " *" + notiType + "* " + " *Notification* \n"
	}

	s += "From: " + noti.Start + " " + "To: " + noti.End + "\n---------------------------\n"
	for _, n := range noti.Tasks {
		s = s + strconv.Itoa(n.ID) + "  " + n.Date + "  " + n.Time + "  " + n.Name + "\n"
	}
	return s
}

func StartNotificationServer() {
	fmt.Println("goline is running...")
	for {
		if ValidateWeek() {
			err := NotifyWeekTasks()
			if err != nil {
				fmt.Println("Something wrong during notify weekly task0", err)
			}
		}
		if ValidateDay() {
			err := NotifyTodayTasks()
			if err != nil {
				fmt.Println("Something wrong during notify daily task", err)
			}
		}
		time.Sleep(1 * time.Second)
	}
}
