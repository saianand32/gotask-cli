package cmd

import (
	"github.com/spf13/cobra"
)

var dropGroupCmd = &cobra.Command{
	Use:   "dropgrp [groupName]",
	Short: "delete a group by name",
	Long:  "delete a group by name",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}
		group := args[0]

		ge := fc.GroupsExecutor(fs)
		if err = ge.DropGroup(group); err != nil {
			return err
		}

		return nil
	},
}
