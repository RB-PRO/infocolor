package gui

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Start() {

	Folder := "infocolor/"

	files, err := ioutil.ReadDir(Folder)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Введите название файла:")
	for _, f := range files {
		fmt.Println(f.Name())
	}

	t := table.NewWriter()
	t.SetTitle("Users")
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle

	t.AppendHeader(table.Row{"Name", "Occupation"})
	t.AppendRow(table.Row{"John Doe", "gardener"})
	t.AppendRow(table.Row{"Roger Roe", "driver"})
	t.AppendRows([]table.Row{{"Paul Smith", "trader"},
		{"Lucy Smith", "teacher"}})

	fmt.Println(t.Render())

}

func Q() {
	fmt.Println("Press 'q' to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" {
			break
		} else {
			fmt.Println("Press 'q' to quit")
		}
	}
}
