package groups

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/saianand32/gotask-cli/internal/filestorage"
	"github.com/saianand32/gotask-cli/internal/helper"
)

type GroupsExecutor interface {
	GetCurrentGroup() (string, error)
	ListGroups() error
	CreateGroup(group string) error
	TruncateGroup(group string) error
	DropGroup(group string) error
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

// ListGroups lists all available groups by scanning the DataFolder for JSON files.
// The current group is highlighted in blue when listed.
func (g *groups) ListGroups() error {
	dirEntries, err := os.ReadDir(g.fs.GetDataFolder())
	if err != nil {
		return fmt.Errorf("couldn't list groups: %v", err)
	}

	currentGroup, _ := g.GetCurrentGroup()
	noGroupsText := ""

	if len(dirEntries) == 0 {
		noGroupsText = "(no groups available)"
	}

	fmt.Println("Available groups:", noGroupsText)
	for _, entry := range dirEntries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".json") {
			groupName := strings.TrimSuffix(entry.Name(), ".json")

			if groupName == currentGroup {
				fmt.Println("- " + helper.Green(groupName))
			} else {
				fmt.Println("- " + groupName)
			}
		}
	}

	return nil
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

func (g *groups) TruncateGroup(group string) error {

	fileName := fmt.Sprintf("%s/%s.json", g.fs.GetDataFolder(), strings.ToLower(group))

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("no group exist named : %v", group)
	}

	data := []byte("[]")
	err := os.WriteFile(fileName, data, 0644)
	if err != nil {
		fmt.Printf("Error truncating file: %v\n", err)
		return err
	}
	fmt.Println("success: Truncated group - ", group)
	return nil
}

func (g *groups) DropGroup(group string) error {

	fileName := fmt.Sprintf("%s/%s.json", g.fs.GetDataFolder(), strings.ToLower(group))

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return fmt.Errorf("no group exist named : %v", group)
	}

	err := os.Remove(fileName)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return err
	}

	cur_group, err := g.GetCurrentGroup()
	if err != nil {
		return fmt.Errorf("fetching current group")
	}

	if strings.EqualFold(cur_group, group) {
		data := []byte("")
		err = os.WriteFile(g.fs.GetGroupFile(), data, 0644)
		if err != nil {
			return fmt.Errorf("couldn't write to file: %v", err)
		}
	}
	fmt.Println("success: Dropped group - ", group)
	return nil
}
