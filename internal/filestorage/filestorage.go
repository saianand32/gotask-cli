package filestorage

import (
	"encoding/json"
	"os"

	"github.com/saianand32/gotask-cli/internal/config"
	"github.com/saianand32/gotask-cli/internal/helper"
	"github.com/saianand32/gotask-cli/internal/models"
)

type FileStorageHandler interface {
	Read(fileName string) ([]models.Item, error)
	Write(fileName string, data []models.Item) error
}

type fs struct {
	BasePath     string
	DataFolder   string
	ConfigFolder string
	GroupFile    string
}

func New(c *config.Config) (*fs, error) {
	if err := helper.InitPaths([]string{c.StoreFolder, c.ConfigFolder, c.DataFolder, c.GroupFile}); err != nil {
		return nil, err
	}
	return &fs{
		BasePath:     c.StoreFolder,
		ConfigFolder: c.ConfigFolder,
		DataFolder:   c.DataFolder,
		GroupFile:    c.GroupFile,
	}, nil
}

func (fs *fs) Read(fileName string) ([]models.Item, error) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []models.Item{}, nil
	}
	// Open the file
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file content
	var items []models.Item
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&items) // Unmarshal JSON into the slice
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (fs *fs) Write(fileName string, data []models.Item) error {
	file, err := os.Create(fileName) // os.Create will either create or truncate the file
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the new data as JSON and write it to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data) // Marshal the data to JSON and write it
	if err != nil {
		return err
	}

	return nil
}
