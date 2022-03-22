package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	. "todocli.mmedic.com/m/v2/src/persistence/db/in_memory_db"
	. "todocli.mmedic.com/m/v2/src/persistence/task_repository"
	. "todocli.mmedic.com/m/v2/src/services/task_service"
)

func main() {
	// -action="action"
	action := flag.String("action", "intro", "Specifies the action for the todo application.")
	flag.Parse()

	PerformAction(*action)

}

func PerformAction(action string) {

	taskDb := CreateTaskDb()
	taskRepository := CreateTaskRepository(taskDb)
	taskService := CreateTaskService(taskRepository)
	errHandler := ErrorHandler()

	switch action {
	case "intro":
		PrintIntro()
	case "add":
		err := taskService.AddTask()
		errHandler(err)
	case "do":
		err := taskService.CompleteTask()
		errHandler(err)
	case "list":
		err := taskService.GetAllTasks()
		errHandler(err)
	case "del":
		err := taskService.DeleteTask()
		errHandler(err)
	default:
		fmt.Printf("Action not defined.\n")
	}
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

func ErrorHandler() func(err error) {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return func(err error) {
		if err != nil {
			errorLog.Fatal(err.Error())
			panic(err)
		}
	}
}
