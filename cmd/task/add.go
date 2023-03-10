package task

import (
	"errors"
	"log"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

var (
	taskName string
	taskFile string = "tasks.json"
)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		throwErr(err)

		if taskName == "" {
			throwErr(errors.New("task name cannot be empty"))
		}

		err = tasks.Add(taskName)
		throwErr(err)

		err = tasks.Save(taskFile)
		throwErr(err)

		log.Println("Task added successfully")
	},
}

func init() {
	AddCmd.Flags().StringVarP(&taskName, "name", "n", "", "The name of the task (Name should be in double quotes)")
	AddCmd.MarkFlagRequired("name")
}

// throwErr is a helper function to throw an error
func throwErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
