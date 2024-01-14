package notification

import (
	"strconv"
	"time"

	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

var (
	cfg config.Config
)

func init() {
	cfg = config.GetSecretConfig()
}

func getWeekDay() (time.Time, time.Time) {
	start := time.Now()

	for start.Weekday() != time.Monday {
		start = start.Add(-time.Hour * 24)
	}

	end := start.Add(time.Hour * 144)
	return start, end
}

func GetTodayTasks() (model.TaskGroup, error) {
	today := time.Now().Format(time.DateOnly)

	noti, err := model.GetByDate(today, today)
	if err != nil {
		return noti, err
	}

	return noti, nil
}

func GetThisWeekTasks() (model.TaskGroup, error) {
	start, end := getWeekDay()

	noti, err := model.GetByDate(start.Format(time.DateOnly), end.Format(time.DateOnly))
	if err != nil {
		return noti, err
	}

	return noti, nil
}

func Format(noti model.TaskGroup) string {
	s := "\n"
	s = s + "From: " + noti.Start + " " + "To: " + noti.End + "\n---------------------------\n"
	for _, n := range noti.Tasks {
		s = s + strconv.Itoa(n.ID) + "  " + n.Date + "  " + n.Time + "  " + n.Name + "\n"
	}
	return s
}
