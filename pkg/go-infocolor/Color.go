package goinfocolor

import "fmt"

// Спарсить цвет и все подобности по этому цвету
func (bz *Brauzer) ParseColor(Link string) (color []ColorForm, err error) {

	return color, err
}

// Подготовить всё о цвете и распечатать
func SprintColorForm(color []ColorForm) (output string) {
	output = "[]ColorForm:" + "\n"
	for iCol, Col := range color {
		output += fmt.Sprintf("%d. Информация о цвете\n", iCol)
		output += fmt.Sprintf("- Марка авто: '%s'\n", Col.Info.Brand)
		output += fmt.Sprintf("- Код краски: '%s'\n", Col.Info.Code)
		output += fmt.Sprintf("- Название цвета: '%s'\n", Col.Info.Name)
		output += fmt.Sprintf("Цвет: '%s'\n", Col.Color)
		output += fmt.Sprintf("Номер панели: '%s'\n", Col.Number)
		output += fmt.Sprintf("Серия: '%s'\n", Col.Seria)
		output += fmt.Sprintf("Покрытие: '%s'\n", Col.Coverage)
		output += fmt.Sprintf("Регион: '%s'\n", Col.Region)
		output += fmt.Sprintf("Оттенок: '%s'\n", Col.Shade)
		output += fmt.Sprintf("Дата раз-ки формулы: '%v'\n", Col.Create.Format("02-01-2006"))
		output += fmt.Sprintf("СТД: '%s'\n", Col.STD)
		output += fmt.Sprintf("Модель: '%s'\n", Col.Model)
		output += fmt.Sprintf("Год выпуска: '%d'\n", Col.Year)
		output += fmt.Sprintf("Производитель: '%s'\n", Col.Manufacturer)
		output += fmt.Sprintf("Дата добавления формулы: '%v'\n", Col.Add.Format("02-01-2006"))
		output += fmt.Sprintf("Автор формулы: '%s'\n", Col.Autor)
		output += ("- Формулы:\n")
		for _, rec := range Col.Rec {
			output += fmt.Sprintf("-- Слой %d:\n", rec.LayerNumber)
			output += fmt.Sprintf("-- Примечание %s:\n", rec.Note)
			output += fmt.Sprintf("-- Сумма %.2f:\n", rec.Coast)
			output += ("-- Элементы формулы:\n")
			for iComp, Comp := range rec.Formuls {
				output += fmt.Sprintf("--- %d. %s\t%s\t%.2f\t%.2f\t\n", iComp, Comp.Code, Comp.Name, Comp.Weight, Comp.CapWeight)
			}
			output += ("-- Элементы комментария:\n")
			for iComm, Comm := range rec.Comments {
				output += fmt.Sprintf("--- %d. %s\t%s\t%s\t\n", iComm, Comm.Data.Format("2 Jan 2006 15:04"), Comm.Autor, Comm.Message)
			}
		}
	}
	return output
}
