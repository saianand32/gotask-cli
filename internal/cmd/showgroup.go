package cmd

import (
	"github.com/spf13/cobra"
)

var showGroupCmd = &cobra.Command{
	Use:   "showgrp",
	Short: "list all groups",
	Long:  "list all groups",
	Args:  cobra.ExactArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fs, err := fc.FileStorage()
		if err != nil {
			return err
		}

		ge := fc.GroupsExecutor(fs)
		if err = ge.ListGroups(); err != nil {
			return err
		}

		return nil
	},
}
