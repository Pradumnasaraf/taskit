package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task        string
	Done        bool
	CreatedOn   string
	CompletedOn string
}

type Tasks []item

func (t *Tasks) Add(task string) {
	Task := item{
		Task:        task,
		Done:        false,
		CreatedOn:   time.Now().Format("02 Jan 2006 15:04:05"),
		CompletedOn: "----",
	}

	*t = append(*t, Task)
}

func (t *Tasks) Complete(taskNum int) error {
	ls := *t // dereference the pointer to the slice of struct
	if taskNum <= 0 || taskNum > len(ls) {
		return errors.New("index out of range")
	}

	index := taskNum - 1 // convert to zero-based index

	ls[index].Done = true
	ls[index].CompletedOn = time.Now().Format("02 Jan 2006 15:04:05")
	return nil

}

func (t *Tasks) Delete(taskNum int) error {
	ls := *t
	if taskNum <= 0 || taskNum > len(ls) {
		return errors.New("index out of range")
	}

	index := taskNum - 1 // convert to zero-based index

	*t = append(ls[:index], ls[index+1:]...)
	return nil
}

func (t *Tasks) Load(filenme string) error {

	file, err := os.ReadFile(filenme)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tasks) Save(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Tasks) Print() {

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "Created On"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	var cells [][]*simpletable.Cell

	for index, items := range *t {
		itemNum := index + 1
		task := blue(items.Task)
		done := blue("No")
		if items.Done {
			task = green(fmt.Sprintf("%s \u2705", items.Task))
			done = green("Yes")
		}

		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", itemNum)},
			{Align: simpletable.AlignLeft, Text: task},
			{Align: simpletable.AlignCenter, Text: done},
			{Align: simpletable.AlignCenter, Text: items.CreatedOn},
			{Align: simpletable.AlignCenter, Text: items.CompletedOn},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending tasks", t.CountPending()))},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())

}

func (t *Tasks) Update(taskNum int, task string) error {
	ls := *t
	if taskNum <= 0 || taskNum > len(ls) {
		return errors.New("index out of range")
	}

	index := taskNum - 1 // convert to zero-based index
	ls[index].Task = task
	return nil
}

func (t *Tasks) DeleteAll() error {

	*t = (*t)[:0]

	return nil
}

func GetInput(arg ...string) (string, error) {
	if len(arg) > 0 {
		return strings.Join(arg, " "), nil
	}

	return "", errors.New("no input provided")
}

func (t *Tasks) CountPending() int {
	var pending int
	for _, item := range *t {
		if !item.Done {
			pending++
		}
	}
	return pending
}
