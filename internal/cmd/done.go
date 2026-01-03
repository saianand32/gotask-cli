package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark task done",
	Long:  "Mark task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}

		ge := fc.GroupsExecutor(fs)
		te := fc.TasksExecutor(fs, ge)
		if err = te.Complete(id); err != nil {
			return err
		}

		fmt.Printf("Task marked done\n")
		return nil
	},
}
