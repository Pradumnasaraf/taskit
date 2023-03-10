package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteAllCmd represents the deleteAll command
var DeleteAllCmd = &cobra.Command{
	Use:   "deleteall",
	Short: "Delete all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteAll called")
	},
}
