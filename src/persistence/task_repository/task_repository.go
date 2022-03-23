package taskrepository

import (
	. "todocli.mmedic.com/m/v2/src/models/task"
	database "todocli.mmedic.com/m/v2/src/persistence/db"
)

type TaskRepository struct {
	db database.DBer
}

func CreateTaskRepository(db database.DBer) *TaskRepository {
	return &TaskRepository{db: db}
}

func (tr *TaskRepository) AddTask(task *Task) {
	tr.db.Add(task)
}

func (tr *TaskRepository) DeleteTask(text string) {
	tr.db.Delete(text)
}

func (tr *TaskRepository) GetTask(text string) *Task {
	return tr.GetTask(text)
}

func (tr *TaskRepository) CompleteTask(text string) {
	task := tr.db.Get(text)
	task.SetStatus(true)
	tr.db.Edit(text, task)
}

func (tr *TaskRepository) GetAllTasks() []*Task {
	return tr.db.GetAll()
}
