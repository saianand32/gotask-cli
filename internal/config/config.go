package config

import (
	"github.com/saianand32/gotask-cli/internal/constants"
)

type Config struct {
	StoreFolder  string
	ConfigFolder string
	DataFolder   string
	GroupFile    string
	Version      string
}

func Default(version string) *Config {
	return &Config{constants.StoreFolder, constants.ConfigFolder, constants.DataFolder, constants.GroupFile, version}
}
