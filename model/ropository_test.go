package model_test

import (
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/tigertony2536/go-line-notify/config"
	"github.com/tigertony2536/go-line-notify/model"
)

func TestGetDB(t *testing.T) {
	cfg := config.GetSecretConfig()
	db := model.GetDB(cfg.DB)
	db.Ping()
	expectType := model.DB{}

	t.Run("get db success", func(t *testing.T) {
		assert.IsTypef(t, &expectType, db, "Expected %T  got %T", &expectType, db)
		assert.NoErrorf(t, db.Ping(), "Connect db successfully")
	})

}

func TestInsertTask(t *testing.T) {
	name := "ส่งคลิปจิตอาสา"
	date := "2023-12-30"
	tm := "00:00:00"
	id, err := model.InsertTask(name, date, tm)
	if err != nil {
		log.Fatal(err)
	}
	expectedNoti, err := model.GetByID(id)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equalf(t, expectedNoti.Tasks[0].Name, name, "Expect %q got %q", expectedNoti.Tasks[0].Name, name)
	assert.Equalf(t, expectedNoti.Tasks[0].Date, date, "Expect %q got %q", expectedNoti.Tasks[0].Date, date)
	assert.Equalf(t, expectedNoti.Tasks[0].Time, tm, "Expect %q got %q", expectedNoti.Tasks[0].Time, tm)
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

		cfg := config.GetSecretConfig()
		db := model.GetDB(cfg.DB)

		noti, err := model.GetByID(tc.ID)

		defer db.Close()

		assert.Equalf(t, tc.ID, noti.Tasks[0].ID, "Expected %q got %q", tc.ID, noti.Tasks[0].ID)
		assert.Equalf(t, tc.ID, noti.Tasks[0].Name, "Expected %q got %q", tc.Message, tc.ID, noti.Tasks[0].Name)
		assert.Equalf(t, tc.ID, noti.Tasks[0].Date, "Expected %q got %q", tc.Date, tc.ID, noti.Tasks[0].Date)
		assert.Equalf(t, tc.ID, noti.Tasks[0].Time, "Expected %q got %q", tc.Time, tc.ID, noti.Tasks[0].Time)
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

func TestGetByName(t *testing.T) {
	tc := []struct {
		Name           string
		pattern        string
		expectedResult model.TaskGroup
	}{
		{
			Name:    "pattern match",
			pattern: "จัด",
			expectedResult: model.TaskGroup{
				Start: "",
				End:   "",
				Tasks: []model.Task{
					{
						ID:   32,
						Name: "จัดทำแผนพัฒนาการศึกษา",
						Date: "2024-02-05",
						Time: "10:00:30",
					},
					{
						ID:   44,
						Name: "ทำงาน1",
						Date: "2023-12-07",
						Time: "09:00:00",
					},
				},
			},
		},
	}
	t.Run(tc[0].Name, func(t *testing.T) {
		group, err := model.GetByName(tc[0].pattern)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, tc[0].expectedResult.Tasks, group.Tasks, "Expected %q got %q", tc[0].expectedResult.Tasks, group.Tasks)
	})
}
