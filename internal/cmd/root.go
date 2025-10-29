package cmd

import (
	"github.com/saianand32/gotask-cli/internal/factory"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gotask",
	Short: "Gotask CLI",
}

func Execute(f factory.Factory) {
	cobra.CheckErr(rootCmd.Execute())
}
