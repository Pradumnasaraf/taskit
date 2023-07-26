package task

import (
	"fmt"
	"log"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

var (
	taskFile string = "tasks.json"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Load the tasks from the file
		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		checkNilErr(err)

		err = tasks.Add(args[0])
		checkNilErr(err)

		err = tasks.Save(taskFile)
		checkNilErr(err)

		fmt.Println("Task added successfully")
	},
}

// checkNilErr is a helper function to throw an error
func checkNilErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
