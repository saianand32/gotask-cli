package factory

import (
	"github.com/saianand32/gotask-cli/internal/components/groups"
	"github.com/saianand32/gotask-cli/internal/components/tasks"
	"github.com/saianand32/gotask-cli/internal/config"
	"github.com/saianand32/gotask-cli/internal/filestorage"
)

type Factory interface {
	FileStorage() (filestorage.FileStorageHandler, error)
	TasksExecutor(fs filestorage.FileStorageHandler, ge groups.GroupsExecutor) tasks.TasksExecutor
	GroupsExecutor(fs filestorage.FileStorageHandler) groups.GroupsExecutor
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

func (f *factory) TasksExecutor(fs filestorage.FileStorageHandler, ge groups.GroupsExecutor) tasks.TasksExecutor {
	return tasks.NewTasksExecutor(fs, ge)
}

func (f *factory) GroupsExecutor(fs filestorage.FileStorageHandler) groups.GroupsExecutor {
	return groups.NewGroupsExecutor(fs)
}
