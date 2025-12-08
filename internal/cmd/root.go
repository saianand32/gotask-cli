package cmd

import (
	"github.com/saianand32/gotask-cli/internal/factory"
	"github.com/spf13/cobra"
)

var (
	rootCmd *cobra.Command
	fc      factory.Factory
)

func Execute(fact factory.Factory) {
	fc = fact
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "gotask",
		Short: "Gotask CLI",
	}

	// Subcommands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(useGroupCmd)
}
