package task

import (
	"errors"
	"log"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

var (
	updateTaskId int
	message      string
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task by ID",
	Run: func(cmd *cobra.Command, args []string) {

		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		throwErr(err)

		if updateTaskId <= 0 {
			throwErr(errors.New("task ID cannot be less than or equal to zero"))
		}

		if message == "" {
			throwErr(errors.New("message cannot be empty"))
		}

		err = tasks.Update(updateTaskId, message)
		throwErr(err)

		err = tasks.Save(taskFile)
		throwErr(err)

		log.Println("Task updated successfully")

	},
}

func init() {
	UpdateCmd.Flags().IntVarP(&updateTaskId, "id", "i", 0, "The ID of the task to be updated")
	UpdateCmd.MarkFlagRequired("id")

	UpdateCmd.Flags().StringVarP(&message, "message", "m", "", "The message to be updated (Message should be in double quotes)")
	UpdateCmd.MarkFlagRequired("message")
}
