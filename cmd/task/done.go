package task

import (
	"errors"
	"log"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

var (
	doneTaskID int
)

// doneCmd represents the done command
var DoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {

		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		throwErr(err)

		if doneTaskID <= 0 {
			throwErr(errors.New("task ID cannot be less than or equal to zero"))
		}

		err = tasks.Done(doneTaskID)
		throwErr(err)

		err = tasks.Save(taskFile)
		throwErr(err)

		log.Println("Task marked as done successfully")

	},
}

func init() {
	DoneCmd.Flags().IntVarP(&doneTaskID, "id", "i", 0, "The ID of the task to be marked as done")
	DoneCmd.MarkFlagRequired("id")
}
