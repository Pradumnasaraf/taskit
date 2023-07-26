package cmd

import (
	"os"

	"github.com/Pradumnasaraf/taskit/cmd/task"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "taskit",
	Short: "A CLI tool to manage your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Add subcommands to root command
	rootCmd.AddCommand(task.AddCmd)
	rootCmd.AddCommand(task.UpdateCmd)
	rootCmd.AddCommand(task.ListCmd)
	rootCmd.AddCommand(task.DeleteCmd)
	rootCmd.AddCommand(task.DeleteAllCmd)
	rootCmd.AddCommand(task.DoneCmd)
}
