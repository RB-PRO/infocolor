package tg

import (
	"fmt"
	"strings"

	minfocolor "github.com/RB-PRO/infocolor/pkg/go-infocolor/m-infocolor"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// // Получить все коды из фломул
// func CodesFile(data []minfocolor.Formulass) (codes []string) {

// 	return RemoveDuplicateStr(codes)
// }

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

// Вытащить все цвета
func ExtractCode(Formulass []minfocolor.Formulass) (Colors []string) {
	for _, formula := range Formulass {
		Colors = append(Colors, formula.PaintCode)
	}
	return RemoveDuplicateStr(Colors)
}

// Вытащить данные по формулам по коду
func ExtractFormulaFromCode(Formulass []minfocolor.Formulass, Code string) (data []minfocolor.Formulass) {
	for _, formula := range Formulass {
		if formula.PaintCode == Code {
			data = append(data, formula)
		}
	}
	return data
}

func PrintFormul(f minfocolor.Formulass) string {
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

	str += "Аттрибуты:\n"
	for iField, Field := range f.Fields {
		str += PrintRow(fmt.Sprintf("%d. %s", iField, Field.Text), Field.Value)
	}

	if len(f.Components) != 0 {
		table := "" // str += TableComponent(f.Components)
		str += "Компоненты формулы: Код, название, вес\n"
		for iComponent, Component := range f.Components {
			table += fmt.Sprintf("%d. %s, %s, %s\n", iComponent+1, Component.Code, Component.Name, Component.Weight)
		}
		str += table + "\n"
	}

	if len(f.Components) != 0 {
		table := "" // str += TableComponentsPrice(f.ComponentsPrice)
		str += "Цены: Код, Density, Rate, Цена\n"
		for iComponentPrice, ComponentPrice := range f.ComponentsPrice {
			table += fmt.Sprintf("%d. %v, %v, %v, %v\n", iComponentPrice+1, ComponentPrice.Name, ComponentPrice.Density, ComponentPrice.Rate, ComponentPrice.Price)
		}
		str += table + "\n"
	}

	if len(f.Components) != 0 { // str += TableCommentaries(f.Commentaries)
		table := ""
		str += "Комментарии: Дата, Автор, Содерджание\n"
		for iComment, Comment := range f.Commentaries {
			Comment.Text = strings.ReplaceAll(Comment.Text, "&nbsp;", "")
			table += fmt.Sprintf("%d. %v, %v, %v\n", iComment+1, Comment.Date, Comment.User, Comment.Text)
		}
		str += table + "\n"
	}

	return str
}

func PrintRow(name string, val string) string {
	var str string
	if val != "" {
		str = fmt.Sprintf("%s: %s\n", name, val)
	}
	return str
}

// Получить значение нижнего бара для отправки сообщения
func Menu(strs []string) (key tgbotapi.ReplyKeyboardMarkup) {
	var j, indexRow int
	Cols := 6
	Rows := ListsCountOfColors(len(strs), Cols)

	// fmt.Println(len(strs), "Rows,Cols", Rows, Cols)

	Col := make([][]tgbotapi.KeyboardButton, Rows)
	for i := 0; i < len(strs); i += Cols {
		j += Cols
		if j > len(strs) {
			j = len(strs)
		}
		// Row := []interface{}{}
		Row := make([]tgbotapi.KeyboardButton, Cols)
		for _, pc := range strs[i:j] {
			Row = append(Row, tgbotapi.NewKeyboardButton(pc))
		}

		Col[indexRow] = Row
		indexRow++
	}

	key.Keyboard = Col

	return key
}

// перевести к-во цветов  в к-во страниц в запросе
// Например для 2 в 1, 15 в 2, 100 в 10
func ListsCountOfColors(a int, Cols int) int {
	b := a / Cols
	if a-b*Cols != 0 {
		return b + 1
	}
	return b
}
