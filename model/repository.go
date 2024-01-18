package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tigertony2536/goline/config"
)

var (
	db  *DB
	cfg config.AppConfig
)

func init() {
	cfg = config.GetAppConfig()
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
	End   string
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
	result, err := db.Exec(`INSERT INTO notify(taskname, date, time) VALUES(?,?,?);`, taskName, date, time)
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

func DeleteTask(id int) error {

	result, err := db.Exec(`DELETE FROM notify WHERE id=?;`, id)
	if err != nil {
		return err
	}
	if n, _ := result.RowsAffected(); int(n) != id {
		s := fmt.Sprintf("Task id %v do not exist", id)
		fmt.Printf("%v", s)
	}
	return nil
}

func UpdateTask(taskId int, taskName, date, time string) error {
	group, err := GetByID(taskId)
	if err != nil {
		return err
	}
	task := group.Tasks[0]

	if taskName != "" {
		task.Name = taskName
	}
	if date != "" {
		task.Date = date
	}
	if time != "" {
		task.Time = time
	}

	result, err := db.Exec(`UPDATE notify 
							SET taskname=?, date=?, time=? 
							WHERE id=?;`, task.Name, task.Date, task.Time, task.ID)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err == nil && n > 0 {
		return nil
	} else {
		return err
	}

}

func GetByID(id int) (TaskGroup, error) {
	result := TaskGroup{"", "", nil}

	row := db.QueryRow("SELECT * FROM notify WHERE id=?;", id)

	m := Task{}
	err := row.Scan(&m.ID, &m.Name, &m.Date, &m.Time)
	if err != nil {
		return result, err
	}
	data := []Task{m}
	result.Tasks = data
	return result, nil
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

func GetByName(pattern string) (TaskGroup, error) {
	result := TaskGroup{"", "", nil}
	rows, err := db.Query(`SELECT * FROM notify WHERE taskname like CONCAT('%',?,'%');`, pattern)
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

func GetAllTasks() (TaskGroup, error) {
	result := TaskGroup{"", "", nil}
	rows, err := db.Query(`SELECT * FROM notify;`)
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
