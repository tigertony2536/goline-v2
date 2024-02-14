package notification

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tigertony2536/goline-v2/config"
)

var (
	appCfg config.AppConfig
	secret config.Secret
)

func init() {
	appCfg = config.GetAppConfig()
	secret = config.GetSecretConfig()
}

func SendToLineApi(notiString string) (string, error) {
	v := url.Values{}
	v.Set("message", notiString)
	client := &http.Client{}
	req, err := http.NewRequest("POST", cfg.Url, strings.NewReader(v.Encode()))

	token := "Bearer " + secret.LineToken

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

func NotifyTodayTasks() error {
	day, err := GetTodayTasks()
	if err != nil {
		return err
	}
	notiString := Format(day, "daily")
	if len(day.Tasks) != 0 {
		_, err := SendToLineApi(notiString)
		if err != nil {
			return err
		}
	}
	return nil
}

func NotifyWeekTasks() error {
	week, err := GetThisWeekTasks()
	if err != nil {
		return err
	}
	notiString := Format(week, "weekly")
	if len(week.Tasks) != 0 {
		_, err := SendToLineApi(notiString)
		if err != nil {
			return err
		}
	}
	return nil
}

func ValidateWeek() bool {
	setDate := appCfg.WeeklyNotiDate
	setTime := appCfg.WeeklyNotiTime
	weekDay := time.Now().Weekday().String()
	sd := strings.ToLower(setDate)
	wd := strings.ToLower(weekDay)
	time := time.Now().Format(time.TimeOnly)
	return wd == sd && time == setTime
}

func ValidateDay() bool {
	setTime := appCfg.DailyNotiTime
	time := time.Now().Format(time.TimeOnly)
	return time == setTime
}
