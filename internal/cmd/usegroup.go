package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var useGroupCmd = &cobra.Command{
	Use:   "usegrp [group]",
	Short: "use/create a new group",
	Long:  "Create a new group or use an existing group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		group := args[0]

		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}

		ge := fc.GroupsExecutor(fs)
		if err = ge.CreateGroup(group); err != nil {
			return err
		}

		fmt.Printf("Using: %s\n", group)
		return nil
	},
}
