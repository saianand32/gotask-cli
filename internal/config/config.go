package config

import (
	"os"
	"path/filepath"

	"github.com/saianand32/gotask-cli/internal/constants"
)

type Config struct {
	StoreFolder  string
	ConfigFolder string
	DataFolder   string
	GroupFile    string
	Version      string
}

func Default(version string) (*Config, error) {
	exePath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	baseDir := filepath.Dir(exePath)

	return &Config{
		StoreFolder:  filepath.Join(baseDir, constants.StoreFolder),
		ConfigFolder: filepath.Join(baseDir, constants.ConfigFolder),
		DataFolder:   filepath.Join(baseDir, constants.DataFolder),
		GroupFile:    filepath.Join(baseDir, constants.GroupFile),
		Version:      version,
	}, nil
}
