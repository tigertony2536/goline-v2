package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tigertony2536/go-line-notify/config"
)

var (
	db  *DB
	cfg config.Config
)

func init() {
	cfg = config.GetConfig()
	db = GetDB(cfg.DB)
}

type Task struct {
	ID   int
	Name string
	Date string
	Time string
}

type TaskGroup struct {
	Start string
	Stop  string
	Tasks []Task
}

type DB struct {
	*sql.DB
}

type Repository interface {
	GetbyID(id int)
}

func GetDB(sc string) *DB {
	db, err := sql.Open("sqlite3", sc)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{db}
}

func InsertTask(taskName, date, time string) (int, error) {
	task := Task{Name: taskName, Date: date, Time: time}
	result, err := db.Exec(`INSERT INTO notify(message, date, time) VALUES(?,?,?);`, task.Name, task.Date, task.Time)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	fmt.Println("Insert the task successfully")
	return int(id), nil
}

func UpdateTask(taskId int, taskName, date, time string) (bool, error) {
	task, err := GetByID(taskId)
	if err != nil {
		return false, err
	}
	if taskName != "" {
		task.Name = taskName
	}
	if date != "" {
		task.Date = date
	}
	if time != "" {
		task.Time = time
	}

	result, err := db.Exec(`UPDATE notify SET id=? taskname=? date=? time=? WHERE id=?;`, taskId, task.Name, task.Date, task.Time)
	if err != nil {
		return false, err
	}
	n, err := result.RowsAffected()
	if err == nil && n > 0 {
		return true, nil
	} else {
		return false, err
	}

}

func GetByID(id int) (Task, error) {
	row := db.QueryRow("SELECT * FROM notify WHERE id=?;", id)

	m := Task{}
	err := row.Scan(&m.ID, &m.Name, &m.Date, &m.Time)
	if err != nil {
		return Task{}, err
	}
	return m, nil
}

func GetByDate(start, end string) (TaskGroup, error) {
	result := TaskGroup{start, end, nil}
	rows, err := db.Query("SELECT * FROM notify WHERE date BETWEEN ? AND ?;", start, end)
	if err != nil {
		return result, err
	}

	data := []Task{}
	defer rows.Close()
	for rows.Next() {
		noti := Task{}
		err := rows.Scan(&noti.ID, &noti.Name, &noti.Date, &noti.Time)
		if err != nil {
			return result, err
		}
		data = append(data, noti)
	}
	result.Tasks = data
	return result, nil
}
