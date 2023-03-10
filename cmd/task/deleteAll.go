package task

import (
	"log"

	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

// deleteAllCmd represents the deleteAll command
var DeleteAllCmd = &cobra.Command{
	Use:   "deleteall",
	Short: "Delete all tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		throwErr(err)

		err = tasks.DeleteAll()
		throwErr(err)

		err = tasks.Save(taskFile)
		throwErr(err)

		log.Println("All tasks deleted successfully")
	},
}
