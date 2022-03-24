package taskservice

import (
	"fmt"

	. "todocli.mmedic.com/m/v2/src/models/task"
	taskRepo "todocli.mmedic.com/m/v2/src/persistence"
)

type TaskService struct {
	tr taskRepo.TaskRepositer
}

func CreateTaskService(tr taskRepo.TaskRepositer) *TaskService {
	return &TaskService{tr: tr}
}

func (ts *TaskService) AddTask(text string) error {
	t := new(Task)
	t.SetText(text)
	err := ts.tr.AddTask(t)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) CompleteTask(text string) error {
	err := ts.tr.CompleteTask(text)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) DeleteTask(text string) error {
	err := ts.tr.DeleteTask(text)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) GetAllTasks() ([]*Task, error) {
	t, err := ts.tr.GetAllTasks()
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (ts *TaskService) ListAllTasks() error {
	tasks, err := ts.tr.GetAllTasks()
	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Printf("-----------Task: %s\nStatus: %t\n", t.GetText(), t.GetStatus())
	}

	return nil
}

func (ts *TaskService) GetCompletedTasks() ([]*Task, error) {
	tasks, err := ts.tr.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var completedTasks []*Task
	for _, t := range tasks {
		if !t.GetStatus() {
			completedTasks = append(completedTasks, t)
		}
	}
	return completedTasks, nil
}

func (ts *TaskService) ListCompletedTasks() error {
	tasks, err := ts.GetCompletedTasks()
	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Printf("-----------\nTask: %s\nStatus: %t\n", t.GetText(), t.GetStatus())
	}

	return err
}

func (ts *TaskService) GetTodaysCompletedTasks() ([]*Task, error) {
	return ts.tr.GetTodaysTasks()
}

func (ts *TaskService) ListTodaysCompletedTasks() error {
	tasks, err := ts.GetTodaysCompletedTasks()
	if err != nil {
		return err
	}

	for _, t := range tasks {
		fmt.Printf("-----------\nTask: %s\nStatus: %t\n", t.GetText(), t.GetStatus())
	}

	return err
}
