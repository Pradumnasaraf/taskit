package task

import (
	"fmt"
	"strconv"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update [task ID] [message]",
	Short: "Update a task by ID",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		checkNilErr(err)

		updateTaskId := args[0]
		convertedString, err := strconv.Atoi(updateTaskId)
		if err != nil {
			fmt.Println("Please enter a valid task ID")
			return
		}

		message := args[1]

		err = tasks.Update(convertedString, message)
		checkNilErr(err)

		err = tasks.Save(taskFile)
		checkNilErr(err)

		fmt.Println("Task updated successfully")

	},
}
