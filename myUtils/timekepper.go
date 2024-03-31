package myutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task         string
	Project      string
	State        string
	Comment      string
	Note         string
	Date         time.Time
	HrsMonday    uint8
	HrsTuesday   uint8
	HrsWednesday uint8
	HrsThursday  uint8
	HrsFriday    uint8
	HrsSaturday  uint8
	HrsSunday    uint8
}

type Projects []item

func (p *Projects) Add(project string, task string, comment string) {
	newProject := item{
		Project: project,
		Task:    task,
		Comment: comment,
		Date:    time.Now(),
	}
	*p = append(*p, newProject)
}

func (p *Projects) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	err = json.Unmarshal(file, p)
	if err != nil {
		return err
	}
	return nil
}
func (p *Projects) Store(filename string) error {
	file, err := json.Marshal(p)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, file, 0644)
}
func (p *Projects) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Project"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignRight, Text: "Comment"},
			{Align: simpletable.AlignRight, Text: "Date"},
			{Align: simpletable.AlignRight, Text: "Mon"},
			{Align: simpletable.AlignRight, Text: "Tue"},
			{Align: simpletable.AlignRight, Text: "Wed"},
			{Align: simpletable.AlignRight, Text: "Thu"},
			{Align: simpletable.AlignRight, Text: "Fri"},
			{Align: simpletable.AlignRight, Text: "Sat"},
			{Align: simpletable.AlignRight, Text: "Sun"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *p {
		idx++
		task := item.Task
		project := item.Project
		comment := item.Comment
		date := item.Date.Format("2006-01-02")
		mon := fmt.Sprintf("%d", item.HrsMonday)
		tue := fmt.Sprintf("%d", item.HrsTuesday)
		wed := fmt.Sprintf("%d", item.HrsWednesday)
		thu := fmt.Sprintf("%d", item.HrsThursday)
		fri := fmt.Sprintf("%d", item.HrsFriday)
		sat := fmt.Sprintf("%d", item.HrsSaturday)
		sun := fmt.Sprintf("%d", item.HrsSunday)

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: project},
			{Text: task},
			{Text: comment},
			{Text: date},
			{Text: mon},
			{Text: tue},
			{Text: wed},
			{Text: thu},
			{Text: fri},
			{Text: sat},
			{Text: sun},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	//	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
	//		{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", 0))},
	//	}}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
