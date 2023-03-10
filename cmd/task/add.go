package task

import (
	"fmt"

	"github.com/spf13/cobra"

)

// addCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}
