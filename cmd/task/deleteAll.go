package task

import (
	"fmt"

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
		checkNilErr(err)

		err = tasks.DeleteAll()
		checkNilErr(err)

		err = tasks.Save(taskFile)
		checkNilErr(err)

		fmt.Println("All tasks deleted successfully")
	},
}
