package goinfocolor

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	playwright "github.com/playwright-community/playwright-go"
)

// Спарсить цвет и все подобности по этому цвету
func (bz *Brauzer) ParseColor(link string) (color []ColorForm, err error) {
	// Навигация на страницу link
	if _, ErrGoto := bz.Page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateCommit,
		Timeout:   playwright.Float(9000),
	}); ErrGoto != nil {
		// fmt.Println(fmt.Errorf("page.Goto: %v", ErrGoto))
		return nil, fmt.Errorf("page.Goto: %v", ErrGoto)
	}

	// Ждём загрузку таблицы
	bz.Page.WaitForSelector(`table[class="color-table-details"]`,
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateAttached,
			Timeout: playwright.Float(20000),
		})

	bz.Page.Click("span[id=show-all-formula-btn]")

	// Инфомрация о цвете - ColorInfo
	var Info ColorInfo
	brand, _ := bz.Page.QuerySelector("table[class=color-info]>tbody>tr:nth-child(1)>td:last-of-type")
	code, _ := bz.Page.QuerySelector("table[class=color-info]>tbody>tr:nth-child(2)>td:last-of-type")
	name, _ := bz.Page.QuerySelector("table[class=color-info]>tbody>tr:nth-child(3)>td:last-of-type")
	if brand != nil {
		Info.Brand, _ = brand.TextContent()
		Info.Brand = strings.TrimSpace(Info.Brand)
	}
	if code != nil {
		Info.Code, _ = code.TextContent()
		Info.Code = strings.TrimSpace(Info.Code)
	}
	if name != nil {
		Info.Name, _ = name.TextContent()
		Info.Name = strings.TrimSpace(Info.Name)
	}

	// // Описание каждого цвета, оно же ColorForm, т.е. каждая строка под красной шапкой
	// RowColor, ErrLocator := bz.Page.QuerySelector(`table[class="color-table-details"]>tbody`)
	// if ErrLocator != nil {
	// 	return nil, fmt.Errorf("bz.Page.QuerySelector: ColorForm: %v", ErrLocator)
	// }

	// Мапа названий колонок, где в качестве ключа номер колонки,
	// а значения - значения в колонке
	ColName := make(map[int]string)
	Collumns_th, ErrLocatorColName := bz.Page.QuerySelectorAll(`table[class="color-table-details"]>tbody>tr:first-of-type>th`)
	if ErrLocatorColName != nil {
		return nil, fmt.Errorf("bz.Page.QuerySelectorAll: ColorForm: ColName %v", ErrLocatorColName)
	}
	for iCol, Col := range Collumns_th {
		if Col != nil {
			val, _ := Col.InnerText()
			val = strings.TrimSpace(val)
			if val != "" {
				ColName[iCol+1] = EditStr(val)
			}
		}
	}
	// fmt.Println("ColName", ColName)

	// // Описание каждого цвета, оно же ColorForm, т.е. каждая строка под красной шапкой
	// RowColor, ErrLocator := bz.Page.QuerySelector(`table[class="color-table-details"]>tbody`)
	// if ErrLocator != nil {
	// 	return nil, fmt.Errorf("bz.Page.QuerySelector: ColorForm: %v", ErrLocator)
	// }

	// Цикл по всей талице со всеми данными
	// первая строка tr - колонки
	// далее смотрим по парам строки,
	//	- первая - общие сведения по цвету,
	//	- вторая - все данные по формуле цвета
	for iRow := 2; ; iRow++ {
		var TecalColorForm ColorForm
		TecalColorForm.Info = Info
		SelectorOneRow := fmt.Sprintf(`tr[class=formula-item]:nth-child(%d)>td`, iRow)
		if Row, ErrLocatorRow := bz.Page.QuerySelectorAll(SelectorOneRow); ErrLocatorRow != nil {
			break // return nil, fmt.Errorf("bz.Page.QuerySelectorAll: ColorForm: ColName %v", ErrLocatorColName)
		} else {
			if len(Row) == 0 {
				break
			}
			for iCol, Col := range Row {
				if Col != nil {
					val, _ := Col.InnerText()
					val = strings.TrimSpace(val)
					// fmt.Println(iCol+1, ColName[iCol+1], val)
					switch ColName[iCol+1] {
					case ("цвет"):
						TecalColorForm.Color = val
					case ("номерпанели"):
						TecalColorForm.Number = val
					case ("серия"):
						TecalColorForm.Seria = val
					case ("покрытие"):
						TecalColorForm.Coverage = val
					case ("регион"):
						TecalColorForm.Region = val
					case ("оттенок"):
						TecalColorForm.Shade = val
					case ("датаразкиформулы"):
						TecalColorForm.Create, _ = time.Parse("02.01.2006", val)
					case ("стд"):
						TecalColorForm.STD = val
					case ("модель"):
						TecalColorForm.Model = val
					case ("годвыпуска"):
						year, _ := strconv.Atoi(val)
						TecalColorForm.Year = year
					case ("производитель"):
						TecalColorForm.Manufacturer = val
					case ("датадобавленияформулы"):
						TecalColorForm.Add, _ = time.Parse("02.01.2006", val)
					case ("авторформулы"):
						TecalColorForm.Autor = val
					}
				}
			}
		}

		// Теперь обрабатываем таблицу с формулами
		var recs []Recipe
		// "[class^=coat-of-pain-item]"
		TdItems, ErrLocatorDivItems := bz.Page.QuerySelectorAll(`td[class=formula-item-components-content]`)
		if ErrLocatorDivItems != nil {
			return nil, fmt.Errorf("bz.Page.QuerySelectorAll: ColorForm: DivItems %v", ErrLocatorDivItems)
		}
		for _, TdItem := range TdItems {
			var rec Recipe
			//
			// Комментарии цвета
			var ColorComments []Comment
			Comments, ErrComments := TdItem.QuerySelectorAll(`ul[id^=commentary-list-]>li`)
			if ErrComments != nil {
				return nil, fmt.Errorf("tdItem.QuerySelectorAll: Comments: %v", ErrComments)
			}
			for _, CommentTeg := range Comments {
				head, _ := CommentTeg.QuerySelector(`span[class=commentary-user]`)
				foot, _ := CommentTeg.QuerySelector(`span[class=commentary-text]`)
				if head != nil && foot != nil {
					headStr, _ := head.TextContent()
					footStr, _ := foot.TextContent()
					Autor := ""
					var Data time.Time
					headStrs := strings.Split(headStr, "|")
					if len(headStrs) == 2 {
						Autor = strings.TrimSpace(headStrs[0])
						Data, _ = time.Parse("02.01.2006 15:04:05", strings.TrimSpace(headStrs[1]))
					}
					footStr = strings.TrimSpace(footStr)
					// fmt.Println(Autor, Data, footStr)
					ColorComments = append(ColorComments, Comment{
						Data:    Data,
						Autor:   Autor,
						Message: footStr,
					})
				}

			}
			rec.Comments = ColorComments

			// блок который содержит всю формулу, цену и прочую перду
			DivItems, ErrLocatorDivItems := TdItem.QuerySelectorAll(`div[id^=coat-of-pain-item-]`)
			if ErrLocatorDivItems != nil {
				return nil, fmt.Errorf("bz.Page.QuerySelectorAll: ColorForm: DivItems %v", ErrLocatorDivItems)
			}
			for _, ItemDiv := range DivItems {

				// Получить название колонок в микротаблице
				ColNameItem, ErrColNameItem := bz.collumnNameItem(ItemDiv)
				if ErrColNameItem != nil {
					return nil, fmt.Errorf("bz.collumnNameItem: %v", ErrColNameItem)
				}

				// Заполнение формул цветов
				Tables, ErrTable := ItemDiv.QuerySelectorAll(`table[class^=color-table]`)
				if ErrTable != nil {
					return nil, fmt.Errorf("Item.QuerySelectorAll: Table: %v", ErrTable)
				}
				for iTable, Table := range Tables {
					var Formuls []Formula
					// fmt.Println(Table.GetAttribute("id"))
					rec.LayerNumber = iTable + 1
					if TR, ErrLocatorRow := Table.QuerySelectorAll(`tbody>tr`); ErrLocatorRow != nil {
						break // return nil, fmt.Errorf("bz.Page.QuerySelectorAll: ColorForm: ColName %v", ErrLocatorColName)
					} else {
						if len(SelectorOneRow) == 0 {
							break
						}
						for itr, tr := range TR {
							if itr == 0 {
								continue
							}
							var Formul Formula
							TD, ErrTD := tr.QuerySelectorAll(`td`)
							if ErrTD != nil {
								return nil, fmt.Errorf("Item.QuerySelectorAll: TD: %v", ErrTD)
							}
							for iCol, td := range TD {
								if td != nil {
									val, _ := td.InnerText()
									val = strings.TrimSpace(val)
									// fmt.Printf("%d. %+v\n", iCol, val)
									if val != "" {
										switch ColNameItem[iCol+1] {
										case ("кодкомпонента"):
											vals := strings.Split(val, "\n")
											if len(vals) > 1 {
												val = vals[0]
											}
											val = strings.TrimSpace(val)
											Formul.Code = val
										case ("названиекомпонента"):
											Formul.Name = val
										case ("весгисходныезначения"):
											val = strings.ReplaceAll(val, ",", ".")
											val = strings.TrimSpace(val)
											Formul.Weight, _ = strconv.ParseFloat(val, 64)
										case ("колвовгчтобыполучилсявыбранныйобъем"):
											val = strings.ReplaceAll(val, ",", ".")
											val = strings.TrimSpace(val)
											Formul.CapWeight, _ = strconv.ParseFloat(val, 64)
										case ("примечание"):
											val = strings.TrimSpace(val)
											rec.Note = val
										}
									}
								}
							}
							Formuls = append(Formuls, Formul) // fmt.Printf("%+v\n", Formul)
						}
					}
					rec.Formuls = Formuls
					// Сумма
					if Sum, ErrSum := ItemDiv.QuerySelector(`span[class=sum]`); ErrSum == nil && Sum != nil {
						// return nil, fmt.Errorf("Item.QuerySelector: Sum: %v", ErrSum)
						sumStr, _ := Sum.InnerText()
						sumStr = strings.ReplaceAll(sumStr, ",", ".")
						sumStr = strings.TrimSpace(sumStr)
						rec.Coast, _ = strconv.ParseFloat(sumStr, 64)
					}

					// Сохранение результатов по формуле
					recs = append(recs, rec)
				}
			}

		}

		//

		TecalColorForm.Rec = recs
		color = append(color, TecalColorForm)
	}

	return color, err
}
func (bz *Brauzer) collumnNameItem(Item playwright.ElementHandle) (map[int]string, error) {
	ColNameItem := make(map[int]string)
	Collumns_th, ErrLocatorColName := Item.QuerySelectorAll(`table>tbody>tr:first-of-type>th`)
	if ErrLocatorColName != nil {
		return nil, fmt.Errorf("Item.QuerySelectorAll: ColNameItem: %v", ErrLocatorColName)
	}
	for iCol, Col := range Collumns_th {
		if Col != nil {
			val, _ := Col.InnerText()
			val = strings.TrimSpace(val)
			if val != "" {
				ColNameItem[iCol+1] = EditStr(val)
			}
		}
	}
	return ColNameItem, nil
}

// Редактировать строку, приводя её к стандартному типу
func EditStr(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "_", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, "/", "")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, ",", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	return str
}

// Подготовить всё о цвете и распечатать
func SprintColorForm(color []ColorForm) (output string) {
	output = "[]ColorForm:" + "\n"
	for iCol, Col := range color {
		output += fmt.Sprintf("%d. Информация о цвете\n", iCol)
		output += fmt.Sprintf("-- Марка авто: '%s'\n", Col.Info.Brand)
		output += fmt.Sprintf("-- Код краски: '%s'\n", Col.Info.Code)
		output += fmt.Sprintf("-- Название цвета: '%s'\n", Col.Info.Name)
		output += fmt.Sprintf("-Цвет: '%s'\n", Col.Color)
		output += fmt.Sprintf("-Номер панели: '%s'\n", Col.Number)
		output += fmt.Sprintf("-Серия: '%s'\n", Col.Seria)
		output += fmt.Sprintf("-Покрытие: '%s'\n", Col.Coverage)
		output += fmt.Sprintf("-Регион: '%s'\n", Col.Region)
		output += fmt.Sprintf("-Оттенок: '%s'\n", Col.Shade)
		output += fmt.Sprintf("-Дата раз-ки формулы: '%v'\n", Col.Create.Format("02-01-2006"))
		output += fmt.Sprintf("-СТД: '%s'\n", Col.STD)
		output += fmt.Sprintf("-Модель: '%s'\n", Col.Model)
		output += fmt.Sprintf("-Год выпуска: '%d'\n", Col.Year)
		output += fmt.Sprintf("--Производитель: '%s'\n", Col.Manufacturer)
		output += fmt.Sprintf("-Дата добавления формулы: '%v'\n", Col.Add.Format("02-01-2006"))
		output += fmt.Sprintf("-Автор формулы: '%s'\n", Col.Autor)
		output += ("-- Формулы:\n")
		for _, rec := range Col.Rec {
			output += fmt.Sprintf("--- Слой %d:\n", rec.LayerNumber)
			output += fmt.Sprintf("--- Примечание %s\n", rec.Note)
			output += fmt.Sprintf("--- Сумма %.2f:\n", rec.Coast)
			output += ("--- Элементы формулы:\n")
			for iComp, Comp := range rec.Formuls {
				output += fmt.Sprintf("---- %d. %s\t%s\t%.2f\t%.2f\t\n", iComp, Comp.Code, Comp.Name, Comp.Weight, Comp.CapWeight)
			}
			output += ("--- Элементы комментария:\n")
			for iComm, Comm := range rec.Comments {
				output += fmt.Sprintf("---- %d. %s\t%s\t%s\t\n", iComm, Comm.Data.Format("2 Jan 2006 15:04"), Comm.Autor, Comm.Message)
			}
		}
	}
	return output
}
