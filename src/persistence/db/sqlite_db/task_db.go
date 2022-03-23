package sqlitedb

import (
	"database/sql"

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
	stmt, err := db.Prepare(`CREATE TABLE if not exists task (
		"id" INTEGER
		"text" TEXT,
		"status" BOOLEAN,
		PRIMARY KEY("id" AUTOINCREMENT);`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nil)
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

func (sd *SqlDb) Add(task *Task) {}

func (sd *SqlDb) Delete(id string) {}

func (sd *SqlDb) Get(id string) *Task {
	return nil
}

func (sd *SqlDb) Edit(id string, task *Task) *Task {
	return nil
}

func (sd *SqlDb) GetAll() []*Task {
	return nil
}
