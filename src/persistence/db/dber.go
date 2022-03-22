package db

import (
	. "todocli.mmedic.com/m/v2/src/models/task"
)

type DBer interface {
	Add(task *Task)
	Delete(id string)
	Get(id string) *Task
	Edit(id string, task *Task) *Task
	GetAll() []*Task
}
