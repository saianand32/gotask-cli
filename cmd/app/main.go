package main

import (
	"fmt"
	"os"

	"github.com/saianand32/gotask-cli/internal/cmd"
	"github.com/saianand32/gotask-cli/internal/config"
	"github.com/saianand32/gotask-cli/internal/factory"
)

var (
	version = "1.0.0"
)

func main() {
	configData, err := config.Default(version)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	f := factory.NewFactory(configData)
	cmd.Execute(f)
}
