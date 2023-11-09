package goinfocolor

import (
	"fmt"
	"strconv"
	"strings"

	playwright "github.com/playwright-community/playwright-go"
)

// Спарсить страницу с типом и номером страницы.
//
// Типы страниц:
//
//   - "Официальные"
//
//   - "Уч. центра"
//
//   - "Колористов"
//
//     https://infocolor.ru/formuls/?ROOT_SECTION_ID=3&color_system=Spies+Hecker&company=0&paint_code=&NAME=&official=true&PAGEN_1=3
func (bb *Brauzer) ParsePage(brand string, types string, page int) (CC []Color, ListCount int, err error) {

	link := ""
	brand = strings.ReplaceAll(brand, " ", "+")
	switch types {
	case "Официальные":
		link = fmt.Sprintf("https://infocolor.ru/formuls/?ROOT_SECTION_ID=3&color_system=Spies+Hecker&company=%s&paint_code=&NAME=&official=true&PAGEN_1=%d", brand, page)
	case "Уч. центра":
		link = fmt.Sprintf("https://infocolor.ru/formuls/?ROOT_SECTION_ID=2&color_system=Spies+Hecker&company=%s&paint_code=&NAME=&training_centre=true&PAGEN_1=%d", brand, page)
	case "Колористов":
		link = fmt.Sprintf("https://infocolor.ru/formuls/?ROOT_SECTION_ID=4&color_system=Spies+Hecker&company=%s&paint_code=&NAME=&colorist=true&PAGEN_1=%d", brand, page)
	}

	// Навигация на страницу link
	if _, ErrGoto := bb.Page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateCommit,
		Timeout:   playwright.Float(9000),
	}); ErrGoto != nil {
		fmt.Println(fmt.Errorf("page.Goto: %v", ErrGoto))
		// return nil, fmt.Errorf("page.Goto: %v", ErrGoto)
	}

	// Ждём загрузку таблицы
	bb.Page.WaitForSelector("table[class=color-table]",
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateAttached,
			Timeout: playwright.Float(20000),
		})

	// Парсинг таблицы
	Row, ErrLocator := bb.Page.QuerySelectorAll("table[class=color-table]>tbody>tr")
	if ErrLocator != nil {
		return nil, 0, fmt.Errorf("bb.Page.QuerySelectorAll: %v", ErrLocator)
	}
	for i, tr := range Row {
		if i == 0 {
			continue
		}

		// Парсинг таблицы
		brand, _ := tr.QuerySelector("td:nth-child(1)")
		code, _ := tr.QuerySelector("td:nth-child(2)")
		name, _ := tr.QuerySelector("td:nth-child(3)")
		note, _ := tr.QuerySelector("td:nth-child(4)")
		link, _ := tr.QuerySelector("td:nth-child(5)>a")

		// Получаем значения
		brandStr, _ := brand.TextContent()
		codeStr, _ := code.TextContent()
		nameStr, _ := name.TextContent()
		noteStr, _ := note.TextContent()
		linkStrHref, _ := link.GetAttribute("href")
		linkStrHref = strings.ReplaceAll(linkStrHref, " ", "+")

		CC = append(CC, Color{
			Info: ColorInfo{
				Brand: strings.TrimSpace(brandStr),
				Code:  strings.TrimSpace(codeStr),
				Name:  strings.TrimSpace(nameStr),
			},
			Note: strings.TrimSpace(noteStr),
			Type: types,
			Link: URL + linkStrHref,
		})
	}

	// Подсчёт того, сколько всего страниц
	List, ErrLists := bb.Page.QuerySelector(`li[class="section-item active"]`)
	if ErrLists != nil {
		return nil, 0, fmt.Errorf("bb.Page.QuerySelectorAll: Поиск к-во листов: %v", ErrLists)
	}
	if List != nil {
		CouterListStr, _ := List.InnerText()
		CouterListStr = strings.ReplaceAll(CouterListStr, "Официальные", "")
		CouterListStr = strings.ReplaceAll(CouterListStr, "Уч. центра", "")
		CouterListStr = strings.ReplaceAll(CouterListStr, "Колористов", "")
		CouterListStr = strings.ReplaceAll(CouterListStr, "(", "")
		CouterListStr = strings.ReplaceAll(CouterListStr, ")", "")
		CouterListStr = strings.TrimSpace(CouterListStr)
		ListCountFromStr, ErrAtoi := strconv.Atoi(CouterListStr)
		if ErrAtoi != nil {
			return nil, 0, fmt.Errorf("strconv.Atoi: '%s': %v", CouterListStr, ErrAtoi)
		}
		ListCount = ListCountFromStr
	}

	return CC, ListsCountOfColors(ListCount), err
}

// перевести к-во цветов  в к-во страниц в запросе
// Например для 2 в 1, 15 в 2, 100 в 10
func ListsCountOfColors(a int) int {
	b := a / 10
	if a-b*10 != 0 {
		return b + 1
	}
	return b
}
