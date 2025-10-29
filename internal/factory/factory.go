package factory

import (
	"github.com/saianand32/gotask-cli/internal/config"
	"github.com/saianand32/gotask-cli/internal/filestorage"
)

type Factory interface {
	FileStorage() (filestorage.FileStorageHandler, error)
}

type factory struct {
	config *config.Config
}

func NewFactory(c *config.Config) Factory {
	return &factory{c}
}

func (f *factory) FileStorage() (filestorage.FileStorageHandler, error) {
	return filestorage.New(f.config)
}
