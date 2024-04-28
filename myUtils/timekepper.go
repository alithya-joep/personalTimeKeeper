package myutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"

	"github.com/alexeyco/simpletable"
)

type item struct {
	Task         string
	Project      string
	State        string
	Comment      string
	Note         string
	Date         string
	HrsMonday    uint8
	HrsTuesday   uint8
	HrsWednesday uint8
	HrsThursday  uint8
	HrsFriday    uint8
	HrsSaturday  uint8
	HrsSunday    uint8
}

type Projects []item

func (p *Projects) Add(project string, task string, comment string, date string) {
	newProject := item{
		Project: project,
		Task:    task,
		Comment: comment,
		Date:    date,
	}
	*p = append(*p, newProject)
}
func (p *Projects) Update(index int, day string, newValue int) {

	ls := *p
	if index < 0 || index >= len(ls) {
		fmt.Println("invalid index")
		return
	}
	index = index - 1

	switch day {
	case "mon":
		ls[index].HrsMonday = uint8(newValue)
	case "tue":
		ls[index].HrsTuesday = uint8(newValue)
	case "wed":
		ls[index].HrsWednesday = uint8(newValue)
	case "thu":
		ls[index].HrsThursday = uint8(newValue)
	case "fri":
		ls[index].HrsFriday = uint8(newValue)
	case "sat":
		ls[index].HrsSaturday = uint8(newValue)
	case "sun":
		ls[index].HrsSunday = uint8(newValue)
	default:
		fmt.Println("invalid day")
	}

	*p = ls

}
func (p *Projects) Delete(index int) error {
	ls := *p
	if index < 0 || index >= len(ls) {
		return errors.New("invaid index")
	}
	*p = append(ls[:index-1], ls[index:]...)
	return nil
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
func (p *Projects) PrintTea() {
	// Define some styles
	// You can use lipgloss to style your tables
	re := lipgloss.NewRenderer(os.Stdout)
	baseStyle := re.NewStyle().Padding(0, 1)
	headerStyle := baseStyle.Copy().Foreground(lipgloss.Color("99")).Bold(true)
	//selectedStyle := baseStyle.Copy().Foreground(lipgloss.Color("#01BE85")).Background(lipgloss.Color("#00432F"))
	headers := []string{"#", "project", "task", "comment", "date", "mon", "tue", "wed", "thu", "fri", "sat", "sun"}
	CapitalizeHeaders := func(data []string) []string {
		for i := range data {
			data[i] = strings.ToUpper(data[i])
		}
		return data
	}
	data := [][]string{}
	for idx, item := range *p {
		data = append(data, []string{
			fmt.Sprint(idx + 1),
			item.Project,
			item.Task,
			item.Comment,
			item.Date,
			fmt.Sprint(item.HrsMonday),
			fmt.Sprint(item.HrsTuesday),
			fmt.Sprint(item.HrsWednesday),
			fmt.Sprint(item.HrsThursday),
			fmt.Sprint(item.HrsFriday),
			fmt.Sprint(item.HrsSaturday),
			fmt.Sprint(item.HrsSunday),
		})
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(re.NewStyle().Foreground(lipgloss.Color("238"))).
		Headers(CapitalizeHeaders(headers)...).
		Width(80).
		Rows(data...).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == 0 {
				return headerStyle
			}

			even := row%2 == 0

			if even {
				return baseStyle.Copy().Foreground(lipgloss.Color("245"))
			}
			return baseStyle.Copy().Foreground(lipgloss.Color("252"))
		})
	fmt.Println(t)
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
		date := item.Date
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
