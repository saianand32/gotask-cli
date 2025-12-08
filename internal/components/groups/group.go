package groups

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/saianand32/gotask-cli/internal/filestorage"
)

type GroupsExecutor interface {
	GetCurrentGroup() (string, error)
	CreateGroup(group string) error
}

type groups struct {
	fs filestorage.FileStorageHandler
}

func NewGroupsExecutor(fs filestorage.FileStorageHandler) *groups {
	return &groups{fs}
}

func (g *groups) GetCurrentGroup() (string, error) {
	file, err := os.Open(g.fs.GetGroupFile())
	if err != nil {
		return "", fmt.Errorf("couldn't open groups file: %v", err)
	}
	defer file.Close()

	groupData, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("couldn't read groups file: %v", err)
	}

	group := strings.TrimSpace(string(groupData))
	if group == "" {
		return "", fmt.Errorf("no group selected. use 'usegrp <group_name>' to select a group")
	}
	return group, nil
}

// CreateGroup creates a new group by storing the group name in the GroupFile and
// Creates an empty JSON file for tasks within the DataFolder.
func (g *groups) CreateGroup(group string) error {
	groupName := strings.ToLower(group)
	data := []byte(groupName)
	err := os.WriteFile(g.fs.GetGroupFile(), data, 0644)
	if err != nil {
		return fmt.Errorf("couldn't write to file: %v", err)
	}

	fileName := fmt.Sprintf("%s/%s.json", g.fs.GetDataFolder(), groupName)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		err := os.WriteFile(fileName, []byte("[]"), 0644)
		if err != nil {
			return fmt.Errorf("couldn't create group JSON file: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("error checking group file: %v", err)
	}

	return nil
}
