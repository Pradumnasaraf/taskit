package task

import (
	"errors"
	"log"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

var (
	deleteTaskID int
)

// deleteCmd represents the delete command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		throwErr(err)

		if deleteTaskID <= 0 {
			throwErr(errors.New("task ID cannot be less than or equal to zero"))
		}

		err = tasks.Delete(deleteTaskID)
		throwErr(err)

		err = tasks.Save(taskFile)
		throwErr(err)

		log.Println("Task deleted successfully")
	},
}

func init() {
	DeleteCmd.Flags().IntVarP(&deleteTaskID, "id", "i", 0, "The ID of the task to be deleted")
	DeleteCmd.MarkFlagRequired("id")
}
