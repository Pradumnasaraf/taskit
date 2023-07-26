package task

import (
	"fmt"
	"strconv"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		checkNilErr(err)

		deleteTaskID := args[0]
		convertedString, err := strconv.Atoi(deleteTaskID)
		if err != nil {
			fmt.Println("Please enter a valid task ID")
			return
		}

		err = tasks.Delete(convertedString)
		checkNilErr(err)

		err = tasks.Save(taskFile)
		checkNilErr(err)

		fmt.Println("Task deleted successfully")
	},
}
