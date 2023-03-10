package task

import (
	"github.com/Pradumnasaraf/taskit/handlers"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := &handlers.Tasks{}
		err := tasks.Load(taskFile)
		throwErr(err)

		err = tasks.List()
		throwErr(err)
	},
}
