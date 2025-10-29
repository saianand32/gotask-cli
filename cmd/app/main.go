package main

import (
	"github.com/saianand32/gotask-cli/internal/cmd"
	"github.com/saianand32/gotask-cli/internal/config"
	"github.com/saianand32/gotask-cli/internal/factory"
)

var (
	version = "1.0.0"
)

func main() {
	c := config.Default(version)
	f := factory.NewFactory(c)
	cmd.Execute(f)
}
