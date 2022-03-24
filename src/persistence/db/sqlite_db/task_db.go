package sqlitedb

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"

	. "todocli.mmedic.com/m/v2/src/models/task"
)

type SqlDb struct {
	db *sql.DB
}

func CreateTaskDb() (*SqlDb, error) {
	db, err := ConnectToDb()
	if err != nil {
		return nil, err
	}
	sqlDb := &SqlDb{db: db}
	return sqlDb, nil
}

func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return nil, err
	}
	SetupTables(db)
	return db, err
}

func SetupTables(db *sql.DB) error {
	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS	 task (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		text TEXT UNIQUE,
		status BOOLEAN,
		completed_at TEXT);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	defer stmt.Close()

	return nil
}

func (sd *SqlDb) CloseConnection() error {
	err := sd.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (sd *SqlDb) Add(task *Task) error {
	tx, err := sd.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO task(text, status, completed_at) VALUES(?, ?, ?);`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.GetText(), task.GetStatus(), "")
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (sd *SqlDb) Delete(text string) error {
	tx, err := sd.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM task WHERE text=?;`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(text)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (sd *SqlDb) Get(text string) (*Task, error) {
	stmt, err := sd.db.Prepare(`SELECT text, status from task WHERE text LIKE ?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(text)

	var txt string
	var sts bool
	err = row.Scan(&txt, &sts)
	if err != nil {
		return nil, err
	}

	t := new(Task)
	t.SetText(txt)
	t.SetStatus(sts)

	return t, nil
}

func (sd *SqlDb) Edit(text string, task *Task) error {

	tx, err := sd.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := sd.db.Prepare(`
	UPDATE task
	SET text = ?,
		status = ?,
		completed_at = ?
	WHERE
		text=?;`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer stmt.Close()

	y, m, d := time.Now().Date()
	today := fmt.Sprintf("%d-%s-%d", d, m, y)
	_, err = stmt.Exec(task.GetText(), task.GetStatus(), today, text)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (sd *SqlDb) GetAll() ([]*Task, error) {
	stmt, err := sd.db.Prepare(`SELECT text, status FROM task;`)
	if err != nil {
		return nil, err
	}
	var tasks []*Task = []*Task{}

	rows, err := stmt.Query()
	for rows.Next() {
		var t Task
		var txt string
		var sts bool
		if err := rows.Scan(&txt, &sts); err != nil {
			return nil, err
		}
		t.SetText(txt)
		t.SetStatus(sts)
		tasks = append(tasks, &t)
	}
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return tasks, nil
}

func (sd *SqlDb) GetFromRange() ([]*Task, error) {

	y, m, d := time.Now().Date()
	today := fmt.Sprintf("%d-%s-%d", d, m, y)

	stmt, err := sd.db.Prepare(`
		SELECT
			text,
			status
		FROM
		task
		WHERE
		completed_at LIKE ?;`)
	if err != nil {
		return nil, err
	}

	var tasks []*Task = []*Task{}

	rows, err := stmt.Query(today)
	for rows.Next() {
		var t Task
		var txt string
		var sts bool
		if err := rows.Scan(&txt, &sts); err != nil {
			return nil, err
		}
		t.SetText(txt)
		t.SetStatus(sts)
		tasks = append(tasks, &t)
	}
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return tasks, nil
}
