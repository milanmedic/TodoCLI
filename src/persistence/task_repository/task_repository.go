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

func (tr *TaskRepository) AddTask(task *Task) error {
	err := tr.db.Add(task)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) DeleteTask(text string) error {
	tr.db.Delete(text)
	return nil
}

func (tr *TaskRepository) GetTask(text string) (*Task, error) {
	t, err := tr.db.Get(text)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (tr *TaskRepository) CompleteTask(text string) error {
	task, err := tr.db.Get(text)
	if err != nil {
		return err
	}
	task.SetStatus(true)
	err = tr.db.Edit(text, task)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) GetAllTasks() ([]*Task, error) {
	t, err := tr.db.GetAll()
	return t, err
}
