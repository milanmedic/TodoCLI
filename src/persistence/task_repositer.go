package taskrepository

import (
	. "todocli.mmedic.com/m/v2/src/models/task"
)

type TaskRepositer interface {
	AddTask(task *Task) error
	DeleteTask(text string) error
	GetTask(text string) (*Task, error)
	CompleteTask(text string) error
	GetAllTasks() ([]*Task, error)
	GetTodaysTasks() ([]*Task, error)
}
