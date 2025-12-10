package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Long:  "Add a new task with a description to your task list",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		description := args[0]

		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}

		ge := fc.GroupsExecutor(fs)
		te := fc.TasksExecutor(fs, ge)
		if err = te.Add(description); err != nil {
			return err
		}

		fmt.Printf("Task added: %s\n", description)
		return nil
	},
}
