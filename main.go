package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	taskDb "todocli.mmedic.com/m/v2/src/persistence/db/sqlite_db"
	taskRepo "todocli.mmedic.com/m/v2/src/persistence/task_repository"
	taskService "todocli.mmedic.com/m/v2/src/services/task_service"
)

func main() {
	// -action="action"
	action := flag.String("action", "add", "Specifies the action for the todo application.")
	flag.Parse()

	PerformAction(*action)

}

func PerformAction(action string) {

	errHandler := ErrorHandler()
	tdb, err := taskDb.CreateTaskDb()
	errHandler(err)
	defer Cleanup(tdb, errHandler)
	tr := taskRepo.CreateTaskRepository(tdb)
	ts := taskService.CreateTaskService(tr)

	switch action {
	case "intro":
		PrintIntro()
	case "add":
		AddTask(ts)
	case "do":
		CompleteTask(ts)
	case "list":
		ListTasks(ts)
	case "del":
		DeleteTask(ts)
	default:
		fmt.Printf("Action not defined.\n")
	}
}

func Cleanup(tdb *taskDb.SqlDb, errHandler func(error)) {
	err := tdb.CloseConnection()
	errHandler(err)
}

func ErrorHandler() func(err error) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return func(err error) {
		if err != nil {
			errorLog.Fatal(err.Error())
			panic(err)
		}
	}
}

func AddTask(ts *taskService.TaskService) {
	fmt.Println("Add task: ")
	var text string
	fmt.Scanf("%s", &text)
	ts.AddTask(text)
	return
}

func CompleteTask(ts *taskService.TaskService) {
	fmt.Println("Complete task: ")
	var text string
	fmt.Scanf("%s", &text)
	ts.CompleteTask(text)
}

func ListTasks(ts *taskService.TaskService) {
	fmt.Println("All tasks")
	ts.ListAllTasks()
}

func DeleteTask(ts *taskService.TaskService) {
	fmt.Println("Delete task: ")
	var text string
	fmt.Scanf("%s", &text)
	ts.DeleteTask(text)
}

func PrintIntro() {
	fmt.Printf(`
	Task is a CLI for managing your TODOs!

	Usage:
	task [command]
  
  	Available Commands:
		add         Add a new task to your TODO list
		do          Mark a task on your TODO list as complete
		list        List all of your incomplete tasks
	
  	Use "task [command] --help" for more information about a command.
	`)
}
