package taskrepository

import (
	. "todocli.mmedic.com/m/v2/src/models/task"
)

type TaskRepositer interface {
	AddTask(task *Task)
	DeleteTask(text string)
	GetTask(text string) *Task
	CompleteTask(text string)
	GetAllTasks() []*Task
}
