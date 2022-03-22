package taskservice

import (
	"fmt"

	. "todocli.mmedic.com/m/v2/src/persistence"
)

type TaskService struct {
	tr TaskRepositer
}

func CreateTaskService(tr TaskRepositer) *TaskService {
	return &TaskService{tr: tr}
}

func (ts *TaskService) AddTask() error {
	fmt.Println("Add Task")
	return nil
}

func (ts *TaskService) CompleteTask() error {
	fmt.Println("Complete Task")
	return nil
}

func (ts *TaskService) DeleteTask() error {
	fmt.Println("Delete Task")
	return nil
}

func (ts *TaskService) GetAllTasks() error {
	fmt.Println("Get All Tasks")
	return nil
}

func (ts *TaskService) GetCompletedTasks() {}
