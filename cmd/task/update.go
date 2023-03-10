package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a task by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("update called")
	},
}
