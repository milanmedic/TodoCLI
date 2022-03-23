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

func (ts *TaskService) AddTask(text string) {
	t := new(Task)
	t.SetText(text)
	ts.tr.AddTask(t)
}

func (ts *TaskService) CompleteTask(text string) {
	ts.tr.CompleteTask(text)
}

func (ts *TaskService) DeleteTask(text string) {
	ts.tr.DeleteTask(text)
}

func (ts *TaskService) GetAllTasks() []*Task {
	return ts.tr.GetAllTasks()
}

func (ts *TaskService) ListAllTasks() {
	tasks := ts.tr.GetAllTasks()
	for _, t := range tasks {
		fmt.Printf(`
		-----------
		Task: %s
		Status %t
		\n`, t.GetText(), t.GetStatus())
	}
}

func (ts *TaskService) GetCompletedTasks() []*Task {
	tasks := ts.tr.GetAllTasks()
	var completedTasks []*Task
	for _, t := range tasks {
		if !t.GetStatus() {
			completedTasks = append(completedTasks, t)
		}
	}
	return completedTasks
}

func (ts *TaskService) ListCompletedTasks() {
	tasks := ts.GetCompletedTasks()
	for _, t := range tasks {
		fmt.Printf(`
		-----------
		Task: %s
		Status %t
		\n`, t.GetText(), t.GetStatus())
	}
}
