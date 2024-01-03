package service

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

var (
	cfg config.Config
)

func init() {
	cfg = config.GetConfig()
}

func getWeekDay() (time.Time, time.Time) {
	start := time.Now()

	for start.Weekday() != time.Monday {
		start = start.Add(-time.Hour * 24)
	}

	end := start.Add(time.Hour * 144)
	return start, end
}

func SendToLineApi(noti model.TaskGroup) (string, error) {
	v := url.Values{}
	v.Set("message", Format(noti))
	client := &http.Client{}
	req, err := http.NewRequest("POST", cfg.Url, strings.NewReader(v.Encode()))

	token := "Bearer " + cfg.LineToken

	if err != nil {
		return "", err
	}
	// req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", token)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	respText, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	s := string(respText)
	return s, nil
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

func NotifyTodayTasks() error {
	day, err := GetTodayTasks()
	if err != nil {
		return err
	}
	if len(day.Tasks) != 0 {
		_, err := SendToLineApi(day)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func NotifyWeekTasks() error {
	week, err := GetThisWeekTasks()
	if err != nil {
		return err
	}
	if len(week.Tasks) != 0 {
		_, err := SendToLineApi(week)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
