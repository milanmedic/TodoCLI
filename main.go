package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	taskDb "todocli.mmedic.com/m/v2/src/persistence/db/sqlite_db"
	taskRepo "todocli.mmedic.com/m/v2/src/persistence/task_repository"
	taskService "todocli.mmedic.com/m/v2/src/services/task_service"
)

func main() {
	// -action="action"
	action := flag.String("action", "list", "Specifies the action for the todo application.")
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
		err := AddTask(ts)
		errHandler(err)
	case "do":
		err := CompleteTask(ts)
		errHandler(err)
	case "list":
		err := ListTasks(ts)
		errHandler(err)
	case "del":
		err := DeleteTask(ts)
		errHandler(err)
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

func PrintIntro() {
	fmt.Printf(`
Task is a CLI for managing your TODOs!
Usage:
task [command]

Available Commands:
	add         Add a new task to your TODO list
	do          Mark a task on your TODO list as complete
	list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.`)
}

func AddTask(ts *taskService.TaskService) error {
	fmt.Println("Add task: ")
	text := ScanInput()
	return ts.AddTask(text)
}

func CompleteTask(ts *taskService.TaskService) error {
	fmt.Println("Complete task: ")
	text := ScanInput()
	return ts.CompleteTask(text)
}

func ListTasks(ts *taskService.TaskService) error {
	fmt.Println("All tasks")
	return ts.ListAllTasks()
}

func DeleteTask(ts *taskService.TaskService) error {
	fmt.Println("Delete task: ")
	text := ScanInput()
	return ts.DeleteTask(text)
}

func ScanInput() string {
	reader := bufio.NewReader(os.Stdin)
	txt, _ := reader.ReadString('\n')
	before, _, _ := strings.Cut(txt, "\n")
	return before
}
