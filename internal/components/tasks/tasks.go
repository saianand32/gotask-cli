package tasks

import (
	"fmt"
	"time"

	"github.com/saianand32/gotask-cli/internal/components/groups"
	"github.com/saianand32/gotask-cli/internal/filestorage"
	"github.com/saianand32/gotask-cli/internal/helper"
	"github.com/saianand32/gotask-cli/internal/models"
)

type TasksExecutor interface {
	Add(task string) error
}

type tasks struct {
	fs    filestorage.FileStorageHandler
	ge    groups.GroupsExecutor
	items []models.Item
}

func NewTasksExecutor(fs filestorage.FileStorageHandler, ge groups.GroupsExecutor) *tasks {
	return &tasks{fs, ge, []models.Item{}}
}

// Add appends a new todo item to the Todos slice and writes it to the specified file.
// It takes the FileStorage instance, the group name, and the task description as parameters.
// If a todo with the same group and task already exists, it returns an error.
func (t *tasks) Add(task string) error {

	group, err := t.ge.GetCurrentGroup()
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s.json", t.fs.GetDataFolder(), group)
	data, err := t.fs.Read(fileName)
	if err != nil {
		return err
	}

	t.items = append(t.items, data...)

	id, err := helper.GenerateCryptoID()
	if err != nil {
		return err
	}

	// Check for duplicates
	for _, v := range data {
		if v.Group == group && v.Task == task {
			return fmt.Errorf("previous task with the same name already exists")
		}
	}

	t.items = append(t.items, models.Item{
		Id:          id,
		Group:       group,
		Task:        task,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	})

	if err = t.fs.Write(fileName, t.items); err != nil {
		return err
	}

	return nil
}
