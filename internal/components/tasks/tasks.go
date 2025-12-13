package tasks

import (
	"fmt"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/saianand32/gotask-cli/internal/components/groups"
	"github.com/saianand32/gotask-cli/internal/filestorage"
	"github.com/saianand32/gotask-cli/internal/helper"
	"github.com/saianand32/gotask-cli/internal/models"
)

type TasksExecutor interface {
	Add(task string) error
	Complete(task string) error
	List() error
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

// Complete marks a todo item as completed by setting the Done field to true
// and updating the CompletedAt field with the current time.
// It searches for the todo item using its ID and returns an error if not found.
func (t *tasks) Complete(id string) error {
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

	for i, todo := range t.items {
		if todo.Id == id {
			if todo.Done {
				return fmt.Errorf("todo with id %s already done", id)
			}
			t.items[i].Done = true
			t.items[i].CompletedAt = time.Now()
			return t.fs.Write(fileName, t.items)
		}
	}
	return fmt.Errorf("todo with id %s not found", id)
}

// Print displays the todos in a formatted table.
// It reads the todos from the specified file and appends them to the Todos slice.
// The table includes a header for the group name and columns for UUID, task, completion status,
// and created/completed timestamps. If there are any errors during file reading, it prints the error message.
func (t *tasks) List() error {

	group, err := t.ge.GetCurrentGroup()
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s/%s.json", t.fs.GetDataFolder(), group)
	data, err := t.fs.Read(fileName)
	if err != nil {
		fmt.Println(err)
		return err
	}
	t.items = append(t.items, data...)

	table := simpletable.New()

	groupHeaderCell := &simpletable.Cell{
		Align: simpletable.AlignCenter,
		Text:  fmt.Sprintf("Group: %s", group),
	}
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 6, Text: groupHeaderCell.Text},
		},
	}

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("group: %s", strings.ToUpper(group))},
			{Align: simpletable.AlignCenter, Text: "uuid"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range t.items {
		idx++
		task := helper.Blue(item.Task)
		done := helper.Blue("no")
		circle := "\033[33m●\033[0m" //yellow circle
		completedAt := "pending"
		if item.Done {
			task = helper.Green(item.Task)
			done = helper.Green("yes")
			circle = "\033[32m●\033[0m" //green circle
			completedAt = item.CompletedAt.Format(time.RFC822)
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: circle},
			{Text: item.Id},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: completedAt},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 6, Text: helper.Red(fmt.Sprintf("You have %d pending todos", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
	return nil
}

// CountPending returns the number of todo items in the Todos slice that are not marked as done.
// It iterates through each todo and increments the count for those that are still pending.
func (t *tasks) CountPending() int64 {
	var count int64
	for _, todo := range t.items {
		if !todo.Done {
			count++
		}
	}
	return count
}
