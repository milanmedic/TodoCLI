package db

import (
	. "todocli.mmedic.com/m/v2/src/models/task"
)

type DBer interface {
	Add(task *Task) error
	Delete(id string) error
	Get(id string) (*Task, error)
	Edit(id string, task *Task) error
	GetAll() ([]*Task, error)
}
