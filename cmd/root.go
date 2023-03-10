package cmd

import (
	"os"

	"github.com/Pradumnasaraf/TaskIt/cmd/task"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "taskit",
	Short: "A CLI tool to manage your tasks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(task.AddCmd)
	rootCmd.AddCommand(task.DeleteCmd)
	rootCmd.AddCommand(task.DeleteAllCmd)
	rootCmd.AddCommand(task.ListCmd)
	rootCmd.AddCommand(task.UpdateCmd)

}
