package gui

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	minfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor/m-infocolor"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func Start() {
	Folder := "infocolor/"

	for {
		fmt.Println(FileList(Folder))
		SpaseWait()

		// ::: Файл
		fmt.Print("Введите название файла: ")
		File := "citroen.json"
		File = InputString()
		Formulas, ErrLoad := LoadFile(Folder + File)
		if ErrLoad != nil {
			fmt.Printf("[Ошибка]: Не найден файл %s%s\n", Folder, File)
			continue
		}
		fmt.Print("\n\n\n")

		// ::: Типы цветов
		fmt.Println(TableTypes(Formulas))
		fmt.Print("Введите тип цвета(2, 3 или 4)(если не хотите, то '0'): ")
		Type := 4
		TypeStr := InputString()
		Type, ErrType := strconv.Atoi(TypeStr)
		if ErrType != nil {
			fmt.Printf("Я не распознал тип '%s', поэтому делаю сразу для всех типов\n", TypeStr)
		}
		if Type == 0 {
			fmt.Println("Вы ввели '0', а значит поиск производится сразу во всех типах")
		}
		var FormulasType []minfocolor.Formulass
		if Type == 2 || Type == 3 || Type == 4 {
			FormulasType = ExetFormulaFromType(Formulas, Type)
		} else {
			FormulasType = Formulas
		}

		// ::: Код цвета
		fmt.Println(TablesCodes(ExetCodeData(FormulasType), "Коды цветов"))
		fmt.Print("Введите код цвета: ")
		Code := "EZR"
		Code = InputString()
		if Code == "" {
			PrintFormulae(Formulas)
			continue
		}
		CodeFormulas := ExetFormulaFromCode(FormulasType, Code)
		if len(CodeFormulas) == 0 {
			fmt.Printf("Не найдены формулы с кодом %s\n", Code)
			continue
		}
		fmt.Print("\n\n\n")

		// ::: Цвет цвета :D
		fmt.Println(TablesCodes(ExetCodeNameData(CodeFormulas), "Названия цветов для кода "+Code))
		fmt.Print("Введите название цвета: ")
		Name := "silver"
		Name = InputString()
		if Name == "" {
			PrintFormulae(CodeFormulas)
			continue
		}
		fmt.Print("\n\n\n")

		NameFormulas := ExetFormulaFromName(CodeFormulas, Name)
		if len(NameFormulas) == 0 {
			fmt.Printf("Не найдены формулы с кодом '%s' и названием цвета '%s'\n", Code, Name)
		}
		fmt.Printf("Найдено всего %d цветов\n", len(NameFormulas))
		PrintFormulae(NameFormulas)
		fmt.Print("\n\n\n")
	}
}

// начпечатать всё что есть
func PrintFormulae(Formulass []minfocolor.Formulass) {
	for _, Formula := range Formulass {
		fmt.Println(PrintFormulass(Formula))
		fmt.Print("\n\n")
	}
}

// Вытащить данные по формулам по коду
func ExetFormulaFromCode(Formulass []minfocolor.Formulass, Code string) (data []minfocolor.Formulass) {
	for _, formula := range Formulass {
		if formula.PaintCode == Code {
			data = append(data, formula)
		}
	}
	return data
}

// Вытащить данные по формулам по названию цвета
func ExetFormulaFromName(Formulass []minfocolor.Formulass, Name string) (data []minfocolor.Formulass) {
	for _, formula := range Formulass {
		if formula.Name == Name {
			data = append(data, formula)
		}
	}
	return data
}

// Вытащить данные по формулам по типу формулы
func ExetFormulaFromType(Formulass []minfocolor.Formulass, Type int) (data []minfocolor.Formulass) {
	for _, formula := range Formulass {
		if formula.Type == Type {
			data = append(data, formula)
		}
	}
	return data
}

// Вытащить все цвета
func ExetCodeNameData(Formulass []minfocolor.Formulass) (Colors []string) {
	for _, formula := range Formulass {
		Colors = append(Colors, formula.Name)
	}
	return RemoveDuplicateStr(Colors)
}

// Вывести коды в таблицу
func TablesCodes(PaintCode []string, tytle string) string {
	t := table.NewWriter()
	t.SetTitle(tytle)
	t.SetAutoIndex(true)
	t.Style().Format.Header = text.FormatTitle
	t.AppendHeader(table.Row{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"})
	Cols := 10
	if len(PaintCode) < Cols {
		Cols = len(PaintCode)
	}
	var j int
	for i := 0; i < len(PaintCode); i += Cols {
		j += Cols
		if j > len(PaintCode) {
			j = len(PaintCode)
		}
		Row := []interface{}{}
		for _, pc := range PaintCode[i:j] {
			Row = append(Row, pc)
		}
		t.AppendRow(Row)
	}
	return t.Render()
}

// Вытащить код из всех формул
func ExetCodeData(Formulass []minfocolor.Formulass) (PaintCode []string) {
	for _, formula := range Formulass {
		PaintCode = append(PaintCode, formula.PaintCode)
	}
	return RemoveDuplicateStr(PaintCode)
}

// Удалить дубликаты в слайсе
func RemoveDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// Загрузить данные из файла
func LoadFile(filename string) (data []minfocolor.Formulass, ErrorFile error) {
	// Открыть файл
	jsonFile, ErrorFile := os.Open(filename)
	if ErrorFile != nil {
		return data, ErrorFile
	}
	defer jsonFile.Close()

	// Прочитать файл и получить массив byte
	jsonData, ErrorFile := io.ReadAll(jsonFile)
	if ErrorFile != nil {
		return data, ErrorFile
	}

	// Распарсить массив byte в структуру
	if ErrorFIle := json.Unmarshal(jsonData, &data); ErrorFIle != nil {
		return data, ErrorFile
	}
	return data, ErrorFile
}

// Ввод текста
func InputString() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}

func SpaseWait() {
	fmt.Print("Чтобы продолжить, нажмите на q ...")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		exit := scanner.Text()
		if exit == "q" || exit == "й" {
			break
		} else {
			fmt.Print("Чтобы продолжить, нажмите на q ...")
		}
	}
}
