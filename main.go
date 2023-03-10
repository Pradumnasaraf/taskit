/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/Pradumnasaraf/TaskIt/cmd"

func main() {
	cmd.Execute()
}

/*
package main

import (
	"flag"
	"fmt"
	"os"

	task "github.com/Pradumnasaraf/TaskIt"
)

const (
	taskFile = "task.json"
)

func main() {

	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int("complete", 0, "complete a task")
	delete := flag.Int("delete", 0, "delete a task")
	list := flag.Bool("list", false, "list all tasks")
	deleteAll := flag.Bool("delete-all", false, "delete all tasks")
	update := flag.Int("update", 0, "update a task")
	message := flag.String("message", "", "message for update")

	flag.Usage = help
	flag.Parse()

	tasks := &task.Tasks{}
	err := tasks.Load(taskFile)
	taskErr("Error loading tasks:", err)

	switch {
	case *add:
		task, err := task.GetInput(flag.Args()...)
		taskErr("Error in adding task:", err)

		tasks.Add(task)
		saveJson(tasks)

	case *complete > 0:
		err := tasks.Complete(*complete)
		taskErr("Error in marking task as complete:", err)

		saveJson(tasks)

	case *delete > 0:
		err := tasks.Delete(*delete)
		taskErr("Error in deleting a task:", err)

		saveJson(tasks)

	case *update > 0:

		if *message != "" {
			err := tasks.Update(*update, *message)
			taskErr("Error in updating a task:", err)

			saveJson(tasks)
		} else {
			fmt.Println("Please provide a message to update")
			os.Exit(1)
		}

	case *deleteAll:
		err := tasks.DeleteAll()
		taskErr("Error in deleting all tasks:", err)

		saveJson(tasks)

	case *list:
		tasks.Print()

	default:
		fmt.Println("No command provided")
		flag.Usage()
		os.Exit(1)
	}
}

func help() {
	text := `
 Usage: ./task [options] [arguments]
  Options:
	-add "task"					Add a new task
	-complete [task no]				Complete a task
	-delete [task no]				Delete a task
	-update [task no] -message "message"		Update a task
	-delete-all					Delete all tasks
	-list						List all tasks
	-help						Show this help
`
	fmt.Println(text)
}

// This function is used to save the tasks to a json file
func saveJson(tasks *task.Tasks) {
	err := tasks.Save(taskFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		os.Exit(1)
	}
}

// This is common error handling function for task functions
func taskErr(errMsg string, err error) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}



*/