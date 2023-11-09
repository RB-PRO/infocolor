package goinfocolor

import (
	"fmt"
	"strings"
	"time"

	playwright "github.com/playwright-community/playwright-go"
)

// Получить список брендов со страницы
//
// Типы страниц:
//
//   - "Официальные"
//
//   - "Уч. центра"
//
//   - "Колористов"
func (bb *Brauzer) BrandList(types string) (colors []string, err error) {

	link := ""
	switch types {
	case "Официальные":
		link = "https://infocolor.ru/formuls/?color_system=Spies+Hecker&official=true"
	case "Уч. центра":
		link = "https://infocolor.ru/formuls/?color_system=Spies+Hecker&training_centre=true"
	case "Колористов":
		link = "https://infocolor.ru/formuls/?color_system=Spies+Hecker&colorist=true"
	}

	if _, ErrGoto := bb.Page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateCommit,
		Timeout:   playwright.Float(10000),
	}); ErrGoto != nil {
		return nil, fmt.Errorf("page.Goto: %v", ErrGoto)
	}

	// Ждём загрузку страницы
	bb.Page.WaitForSelector(`select[name=company]>option[value=0]`,
		playwright.PageWaitForSelectorOptions{
			// State:   playwright.WaitForSelectorStateHidden, //WaitForSelectorStateHidden
			Timeout: playwright.Float(10000),
		})
	time.Sleep(time.Second)

	Row, ErrLocator := bb.Page.QuerySelectorAll("select[name=company]>option")
	if ErrLocator != nil {
		return nil, fmt.Errorf("bb.Page.QuerySelectorAll: %v", ErrLocator)
	}

	for _, tr := range Row {
		val, _ := tr.InnerText()
		if !strings.Contains(val, "Выберите марку автомобиля") {
			colors = append(colors, val)
		}
	}
	return colors, nil
}
