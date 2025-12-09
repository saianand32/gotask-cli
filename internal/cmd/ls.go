package cmd

import (
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all tasks in current group.",
	Long:  "List all tasks in current group.",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}

		ge := fc.GroupsExecutor(fs)
		te := fc.TasksExecutor(fs, ge)
		if err = te.List(); err != nil {
			return err
		}

		return nil
	},
}
