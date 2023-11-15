package gui

import (
	"fmt"
	"log"
	"os"
	"strings"

	minfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor/m-infocolor"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

// Вывести все файлы
func FileList(Folder string) string {
	files, err := os.ReadDir(Folder)
	if err != nil {
		log.Fatal(err)
	}
	FileNames := make([]string, 0, 11)
	for _, f := range files {
		FileNames = append(FileNames, f.Name())
	}

	t := table.NewWriter()
	t.SetTitle("Файлы")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(table.Row{"A", "B", "C", "D", "E"})
	Cols := 5
	for i := 0; i < len(FileNames); i += Cols {
		// fmt.Println(len(FileNames[i:i+Cols]), FileNames[i:i+Cols])
		Row := []interface{}{}
		for j := 0; j < len(FileNames[i:i+Cols]); j++ {
			Row = append(Row, FileNames[i : i+Cols][j])
		}
		t.AppendRow(Row)
	}
	return t.Render() + "\n"
}

// Сделать таблицу из компонентов
func TableComponent(Component []minfocolor.Components) string {
	if len(Component) == 0 {
		return ""
	}
	t := table.NewWriter()
	t.SetTitle("Компоненты формулы")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"Code", "CodeDisplay", "Name", "Weight", "WeightDisplay", "WeightDisplayInit"})
	for _, c := range Component {
		t.AppendRow(table.Row{c.Code, c.CodeDisplay, c.Name, c.Weight, c.WeightDisplay, c.WeightDisplayInit})
	}
	return t.Render() + "\n"
}

// Сделать таблицу из компонентов аттрибустов
func TableFields(Component []minfocolor.Fields) string {
	if len(Component) == 0 {
		return ""
	}
	t := table.NewWriter()
	t.SetTitle("Аттрибуты")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(table.Row{"Аттрибут", "Значение"})
	for _, c := range Component {
		t.AppendRow(table.Row{c.Text, c.Value})
	}
	return t.Render() + "\n"
}

// Сделать таблицу из типов цветов и указать их к-во
func TableTypes(forms []minfocolor.Formulass) string {
	if len(forms) == 0 {
		return ""
	}
	t := table.NewWriter()
	t.SetTitle("Типы цветов")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"2", "3", "4"})
	Types := []int{2, 3, 4}
	var Row []interface{}
	for _, Type := range Types {
		var count int
		for _, form := range forms {
			if form.Type == Type {
				count++
			}
		}
		var RowStr string
		switch Type {
		case 2:
			RowStr = fmt.Sprintf("Официальный(%d)", count)
		case 3:
			RowStr = fmt.Sprintf("Уч. цвета(%d)", count)
		case 4:
			RowStr = fmt.Sprintf("Колористов(%d)", count)
		}
		Row = append(Row, RowStr)
	}
	t.AppendRow(Row)
	return t.Render() + "\n"
}

// Сделать таблицу из компонентов цен
func TableComponentsPrice(Component []minfocolor.ComponentsPrice) string {
	if len(Component) == 0 {
		return ""
	}
	t := table.NewWriter()
	t.SetTitle("Компоненты цены")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"Name", "Currency", "Density", "Rate", "Price"})
	for _, c := range Component {
		t.AppendRow(table.Row{c.Name, c.Currency, c.Density, c.Rate, c.Price})
	}
	return t.Render() + "\n"
}

// Сделать таблицу из компонентов цен
func TableCommentaries(Component []minfocolor.Commentaries) string {
	if len(Component) == 0 {
		return ""
	}
	t := table.NewWriter()
	t.SetTitle("Комментарии")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.SetStyle(table.StyleRounded)
	t.AppendHeader(table.Row{"Date", "User", "Text"})
	for _, c := range Component {
		t.AppendRow(table.Row{c.Date, c.User, strings.ReplaceAll(c.Text, "&nbsp;", "")})
	}
	return t.Render() + "\n"
}
func PrintFormulass(f minfocolor.Formulass) string {
	var str string
	switch f.Type {
	case 2:
		str = "Тип: Официальный\n"
	case 3:
		str = "Тип: Уч. цвета\n"
	case 4:
		str = "Тип: Колористов\n"
	}

	str += PrintRow("ID", f.ID)
	str += PrintRow("Company", f.Company)
	str += PrintRow("PaintCode", f.PaintCode)
	str += PrintRow("Name", f.Name)
	str += PrintRow("Quantity", f.Quantity)
	str += PrintRow("Thinning", f.Thinning)
	str += PrintRow("Unit", f.Unit)
	str += PrintRow("UnitSource", f.UnitSource)
	str += PrintRow("Price", f.Price)
	str += PrintRow("PriceLabel", f.PriceLabel)
	str += PrintRow("Sum", f.Sum)

	str += TableFields(f.Fields)
	str += TableComponent(f.Components)
	str += TableComponentsPrice(f.ComponentsPrice)
	str += TableCommentaries(f.Commentaries)

	return str
}

func PrintRow(name string, val string) string {
	var str string
	if val != "" {
		str = fmt.Sprintf("%s: %s\n", name, val)
	}
	return str
}
