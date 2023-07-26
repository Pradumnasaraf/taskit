package task

import (
	"fmt"
	"strconv"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var DoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		checkNilErr(err)

		doneTaskID := args[0]
		convertedString, err := strconv.Atoi(doneTaskID)
		if err != nil {
			fmt.Println("Please enter a valid task ID")
			return
		}

		err = tasks.Done(convertedString)
		checkNilErr(err)

		err = tasks.Save(taskFile)
		checkNilErr(err)

		fmt.Println("Task marked as done successfully")

	},
}
