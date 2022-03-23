package inmemorydb

import (
	. "todocli.mmedic.com/m/v2/src/models/task"
)

type TaskInMemDb struct {
	tasks map[string]*Task
}

var taskDb *TaskInMemDb = nil

func CreateTaskDb() *TaskInMemDb {
	if taskDb == nil {
		taskDb = &TaskInMemDb{tasks: make(map[string]*Task)}
	}
	return taskDb
}

func (tm *TaskInMemDb) Add(task *Task) {
	tm.tasks[task.GetText()] = task
}

func (tm *TaskInMemDb) Delete(id string) {
	if _, ok := tm.tasks[id]; ok {
		delete(taskDb.tasks, id)
	}
}

func (tm *TaskInMemDb) Get(id string) *Task {
	return tm.tasks[id]
}

func (tm *TaskInMemDb) Edit(id string, task *Task) *Task {
	if _, ok := tm.tasks[id]; ok {
		tm.tasks[id] = task
	}
	return nil
}

func (tm *TaskInMemDb) GetAll() []*Task {
	var tasks []*Task
	for _, v := range tm.tasks {
		tasks = append(tasks, v)
	}
	return tasks
}
