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

	// --- Subcommands ---

	// - group commands -
	rootCmd.AddCommand(useGroupCmd)
	rootCmd.AddCommand(showGroupCmd)

	// - task commands -
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(lsCmd)
}
