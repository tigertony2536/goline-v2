package model_test

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tigertony2536/goline/config"
	"github.com/tigertony2536/goline/model"
)

var (
	db  *model.DB
	cfg config.AppConfig
)

func init() {
	cfg = config.GetAppConfig()
	db = model.GetDB(cfg.DB)
}

func TestGetDB(t *testing.T) {
	err := db.Ping()

	if err != nil {
		t.Error("Can not access database")
	}

}

func TestInsertTask(t *testing.T) {
	expect := model.Task{
		Name: "ส่งคลิปจิตอาสา",
		Date: "2023-12-30",
		Time: "00:00:00",
	}

	id, err := model.InsertTask(expect.Name, expect.Date, expect.Time)
	if err != nil {
		log.Fatal("Something wrong during insert task", err)
	}
	fmt.Println(id)

	defer model.DeleteTask(id)

	task, err := model.GetByID(id)
	if err != nil {
		log.Fatal("Something wrong during get task", err)
	}
	expect.ID = id

	fmt.Println(expect)
	fmt.Println(task.Tasks[0])
	if !reflect.DeepEqual(expect, task.Tasks[0]) {
		t.Error("insert task fail")
	}

}

func TestGetByID(t *testing.T) {
	expect := model.Task{
		Name: "ส่งคลิปจิตอาสา",
		Date: "2023-12-30",
		Time: "22:48:52",
	}

	id, err := model.InsertTask(expect.Name, expect.Date, expect.Time)
	if err != nil {
		log.Fatal("Something wrong during insert task", err)
	}

	expect.ID = id

	defer model.DeleteTask(id)

	t.Run("get by id", func(t *testing.T) {

		noti, err := model.GetByID(id)
		if err != nil {
			log.Fatal("Something wrong during get task", err)
		}

		fmt.Println(expect)
		fmt.Println(noti.Tasks[0])
		if !reflect.DeepEqual(expect, noti.Tasks[0]) {
			t.Error("Get by id fail")
		}
	})
}

func TestGetByDate(t *testing.T) {
	tc := []model.Task{
		{Name: "taskA", Date: "3000-01-01", Time: "99-99-01"},
		{Name: "taskA", Date: "3000-01-02", Time: "99-99-02"},
		{Name: "taskA", Date: "3000-01-03", Time: "99-99-03"},
		{Name: "taskA", Date: "3000-01-04", Time: "99-99-04"},
	}

	ids := []int{}
	for index, tsk := range tc {
		id, err := model.InsertTask(tsk.Name, tsk.Date, tsk.Time)
		if err != nil {
			log.Fatal("Something wrong during insert task", err)
		}
		tc[index].ID = id
		ids = append(ids, id)
	}

	del := func() {
		for _, id := range ids {
			model.DeleteTask(id)
		}
	}

	defer del()

	t.Run(tc[0].Name, func(t *testing.T) {
		noti, err := model.GetByDate(tc[1].Date, tc[2].Date)
		if err != nil {
			log.Fatal("Get by date fail", err)
		}
		fmt.Println(noti.Tasks[0])
		fmt.Println(tc[1])
		if !reflect.DeepEqual(noti.Tasks[0], tc[1]) {
			t.Error("get by date fail: first task wrong")
		}
		fmt.Println(noti.Tasks[1])
		fmt.Println(tc[2])
		if !reflect.DeepEqual(noti.Tasks[1], tc[2]) {
			t.Error("get by date fail: second task wrong")
		}
	})
}

func TestGetByName(t *testing.T) {
	expect := model.Task{
		Name: "ไก่จิกเด็กตายบนปากโอ่ง",
		Date: "2023-12-30",
		Time: "22:48:52",
	}

	id, err := model.InsertTask(expect.Name, expect.Date, expect.Time)
	if err != nil {
		log.Fatal(err)
	}

	expect.ID = id

	defer model.DeleteTask(id)

	t.Run("get by name", func(t *testing.T) {
		group, err := model.GetByName("บนปาก")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("expect:", expect)
		fmt.Println("result:", group.Tasks[0])
		if !reflect.DeepEqual(expect, group.Tasks[0]) {
			t.Error("get by name fail")
		}
	})
}

func TestUpdateTask(t *testing.T) {
	task := model.Task{
		Name: "task1",
		Date: "2024-01-01",
		Time: "12:00:00",
	}
	id, err := model.InsertTask(task.Name, task.Date, task.Time)
	if err != nil {
		fmt.Println("Something wrong during insert task", err)
	}

	task.ID = id
	defer model.DeleteTask(id)

	expect := []model.Task{
		{id, "task2", "2024-01-01", "12:00:00"},
		{id, "task1", "2024-01-30", "12:00:00"},
		{id, "task1", "2024-01-01", "23:00:00"},
	}
	err = model.UpdateTask(id, expect[0].Name, expect[0].Date, expect[0].Time)
	if err != nil {
		fmt.Println("Update Name fail: ", err)
	}

	result, err := model.GetByID(id)
	if err != nil {
		fmt.Println("Something wrong during get task", err)
	}

	t.Run("update name", func(t *testing.T) {

		fmt.Println("expect:", expect[0])
		fmt.Println("result:", result.Tasks[0])
		if !reflect.DeepEqual(expect[0], result.Tasks[0]) {
			t.Error("update name fail")
		}
	})
}
