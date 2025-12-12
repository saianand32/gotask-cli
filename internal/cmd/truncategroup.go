package cmd

import (
	"github.com/spf13/cobra"
)

var truncateGroupCmd = &cobra.Command{
	Use:   "truncategrp [groupName]",
	Short: "truncate group by name",
	Long:  "clear a group by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}
		group := args[0]

		ge := fc.GroupsExecutor(fs)
		if err = ge.TruncateGroup(group); err != nil {
			return err
		}

		return nil
	},
}
