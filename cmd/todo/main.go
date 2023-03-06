package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/Pradumnasaraf/Go-todo-CLI"
)

const (
	todoFile = "todo.json"
)

func main() {

	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "complete a todo")
	delete := flag.Int("delete", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")
	deleteAll := flag.Bool("delete-all", false, "delete all todos")
	update := flag.Int("update", 0, "update a todo")
	message := flag.String("message", "", "message for update")

	flag.Usage = help
	flag.Parse()

	todos := &todo.Todos{} // Here we are creating a pointer to a slice of struct
	err := todos.Load(todoFile)
	taskErr("Error loading todos:", err)

	switch {
	case *add:
		task, err := todo.GetInput(flag.Args()...)
		taskErr("Error in adding todo:", err)

		todos.Add(task)
		saveJson(todos)

	case *complete > 0:
		err := todos.Complete(*complete)
		taskErr("Error in marking todo as complete:", err)

		saveJson(todos)

	case *delete > 0:
		err := todos.Delete(*delete)
		taskErr("Error in deleting a todo:", err)

		saveJson(todos)

	case *update > 0:
		if *message != "" {
			err := todos.Update(*update, *message)
			taskErr("Error in updating a todo:", err)

			saveJson(todos)
		} else {
			fmt.Println("Please provide a message to update")
			os.Exit(1)
		}

	case *deleteAll:
		err := todos.DeleteAll()
		taskErr("Error in deleting all todos:", err)

		saveJson(todos)

	case *list:
		todos.Print()

	default:
		fmt.Println("No command provided")
		flag.Usage()	
		os.Exit(1)
	}
}

func help() {
	text := `
 Usage: ./todo [options] [arguments]
  Options:
	-add "task"					Add a new todo
	-complete [todo no]				Complete a todo
	-delete [todo no]				Delete a todo
	-update [todo no] -message "message"		Update a todo
	-delete-all					Delete all todos
	-list						List all todos
	-help						Show this help
`
	fmt.Println(text)
}

// This function is used to save the todos to a json file
func saveJson(todos *todo.Todos) {
	err := todos.Save(todoFile)
	if err != nil {
		fmt.Println("Error saving todos:", err)
		os.Exit(1)
	}
}

// This is common error handling function for todo functions
func taskErr(errMsg string, err error) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}
