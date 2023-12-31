package model_test

import (
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

func TestGetDB(t *testing.T) {
	cfg := config.GetConfig()
	db := model.GetDB(cfg.DB)
	db.Ping()
	expectType := model.DB{}

	t.Run("get db success", func(t *testing.T) {
		assert.IsTypef(t, &expectType, db, "Expected %T  got %T", &expectType, db)
		assert.NoErrorf(t, db.Ping(), "Connect db successfully")
	})

}

func TestInsertNotification(t *testing.T) {
	tm := time.Now().Format(time.TimeOnly)

	name := "ส่งคลิปจิตอาสา"
	date := "2023-12-30"

	id, err := model.InsertTask(name, date, tm)

	expectedNoti, _ := model.GetByID(id)

	assert.Equalf(t, expectedNoti.Name, name, "Expect %q got %q", expectedNoti.Name, name)
	assert.Equalf(t, expectedNoti.Date, date, "Expect %q got %q", expectedNoti.Date, date)
	assert.Equalf(t, expectedNoti.Time, tm, "Expect %q got %q", expectedNoti.Time, tm)
	assert.NoError(t, err, "Insert notification to database successfully")

}

func TestGetByID(t *testing.T) {
	tc := struct {
		Name    string
		ID      int
		Message string
		Date    string
		Time    string
	}{
		Name:    "Get by ID Success",
		ID:      29,
		Message: "ส่งคลิปจิตอาสา",
		Date:    "2023-12-30",
		Time:    "22:48:52",
	}

	t.Run(tc.Name, func(t *testing.T) {

		cfg := config.GetConfig()
		db := model.GetDB(cfg.DB)

		noti, err := model.GetByID(tc.ID)

		defer db.Close()

		assert.Equalf(t, tc.ID, noti.ID, "Expected %q got %q", tc.ID, noti.ID)
		assert.Equalf(t, tc.Message, noti.Name, "Expected %q got %q", tc.Message, noti.Name)
		assert.Equalf(t, tc.Date, noti.Date, "Expected %q got %q", tc.Date, noti.Date)
		assert.Equalf(t, tc.Time, noti.Time, "Expected %q got %q", tc.Time, noti.Time)
		assert.NoError(t, err, "No error")
	})

}

func TestGetByDate(t *testing.T) {
	tc := []struct {
		Name               string
		Start              string
		End                string
		ExpectedRowsNumber int
		ExpectedNotiID     []int
	}{
		{
			Name:               "Get by date success: multiple noti",
			Start:              "2023-12-01",
			End:                "2023-12-07",
			ExpectedRowsNumber: 2,
			ExpectedNotiID:     []int{31, 44},
		},
	}

	t.Run(tc[0].Name, func(t *testing.T) {
		noti, err := model.GetByDate(tc[0].Start, tc[0].End)

		notiID := []int{}

		for _, n := range noti.Tasks {
			notiID = append(notiID, n.ID)
		}
		assert.Equalf(t, tc[0].ExpectedRowsNumber, len(noti.Tasks), "Expect %d of  result got %d", tc[0].ExpectedRowsNumber, len(noti.Tasks))
		assert.Equalf(t, tc[0].ExpectedNotiID, notiID, "Expect result ID %d got %d", tc[0].ExpectedNotiID, notiID)
		assert.NoError(t, err, "No Error")
	})
}
