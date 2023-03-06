package todo

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

type Todos []item

func (t *Todos) Add(task string) { //  ()
	todo := item{
		Task:        task,
		Done:        false,
		CreatedOn:   time.Now().Format("02 Jan 2006 15:04:05"),
		CompletedOn: "----",
	}

	*t = append(*t, todo)
}

func (t *Todos) Complete(todoNum int) error {
	ls := *t // dereference the pointer to the slice of struct
	if todoNum <= 0 || todoNum > len(ls) {
		return errors.New("index out of range")
	}

	index := todoNum - 1 // convert to zero-based index

	ls[index].Done = true
	ls[index].CompletedOn = time.Now().Format("02 Jan 2006 15:04:05")
	return nil

}

func (t *Todos) Delete(todoNum int) error {
	ls := *t
	if todoNum <= 0 || todoNum > len(ls) {
		return errors.New("index out of range")
	}

	index := todoNum - 1 // convert to zero-based index

	*t = append(ls[:index], ls[index+1:]...)
	return nil
}

func (t *Todos) Load(filenme string) error {

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

func (t *Todos) Save(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {

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

func GetInput(arg ...string) (string, error) {
	if len(arg) > 0 {
		return strings.Join(arg, " "), nil
	}

	return "", errors.New("no input provided")
}

func (t *Todos) CountPending() int {
	var pending int
	for _, item := range *t {
		if !item.Done {
			pending++
		}
	}
	return pending
}
