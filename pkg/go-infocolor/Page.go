package goinfocolor

import (
	"fmt"
	"strconv"
	"strings"

	playwright "github.com/playwright-community/playwright-go"
)

// 9163
type Brauzer struct {
	PW      *playwright.Playwright
	Browser playwright.Browser
	Page    playwright.Page
}

func NewBrauzer() (*Brauzer, error) {
	// ErrInstall := playwright.Install()
	// if ErrInstall != nil {
	// 	return nil, fmt.Errorf("playwright.Install: %v", ErrInstall)
	// }
	pw, ErrRun := playwright.Run()
	if ErrRun != nil {
		return nil, fmt.Errorf("playwright.Run: %v", ErrRun)
	}
	browser, ErrLaunch := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if ErrLaunch != nil {
		return nil, fmt.Errorf("pw.Chromium.Launch: %v", ErrLaunch)
	}
	pp, ErrNewPage := browser.NewPage()
	if ErrNewPage != nil {
		return nil, fmt.Errorf("browser.NewPage: %v", ErrNewPage)
	}
	return &Brauzer{
		PW:      pw,
		Browser: browser,
		Page:    pp,
	}, nil
}
func (bb *Brauzer) Close() {
	bb.Browser.Close()
	bb.PW.Stop()
}

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

	if _, ErrGoto := bb.Page.Goto(link, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateCommit,
		Timeout:   playwright.Float(3000),
	}); ErrGoto != nil {
		fmt.Println(fmt.Errorf("page.Goto: %v", ErrGoto))
		// return nil, fmt.Errorf("page.Goto: %v", ErrGoto)
	}

	// Ждём загрузку таблицы
	bb.Page.WaitForSelector("table[class=color-table]",
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateAttached,
			Timeout: playwright.Float(15000),
		})

	// Раздел с таблицей
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

		CC = append(CC, Color{
			Info: ColorInfo{
				Brand: strings.TrimSpace(brandStr),
				Code:  strings.TrimSpace(codeStr),
				Name:  strings.TrimSpace(nameStr),
			},
			Note: strings.TrimSpace(noteStr),
			Link: linkStrHref,
		})
	}

	// Подсчёт того, сколько всего страниц
	List, ErrLists := bb.Page.QuerySelector(`li[class="section-item active"]`)
	if ErrLists != nil {
		return nil, 0, fmt.Errorf("bb.Page.QuerySelectorAll: Поиск к-во листов: %v", ErrLists)
	}
	CouterListStr, _ := List.InnerText()
	CouterListStr = strings.ReplaceAll(CouterListStr, "Официальные", "")
	CouterListStr = strings.ReplaceAll(CouterListStr, "Уч. центра", "")
	CouterListStr = strings.ReplaceAll(CouterListStr, "Колористов", "")
	CouterListStr = strings.ReplaceAll(CouterListStr, "(", "")
	CouterListStr = strings.ReplaceAll(CouterListStr, ")", "")
	CouterListStr = strings.TrimSpace(CouterListStr)
	ListCount, ErrAtoi := strconv.Atoi(CouterListStr)
	if ErrAtoi != nil {
		return nil, 0, fmt.Errorf("strconv.Atoi: '%s': %v", CouterListStr, ErrAtoi)
	}

	// Ждём появление селектора с таблицей цветов
	// pp.WaitForSelector("table[class=color-table]")

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
		Timeout:   playwright.Float(30000),
	}); ErrGoto != nil {
		fmt.Println(fmt.Errorf("page.Goto: %v", ErrGoto))
		// return nil, fmt.Errorf("page.Goto: %v", ErrGoto)
	}

	// Ждём загрузку страницы
	bb.Page.WaitForSelector("select[name=company]",
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateHidden,
			Timeout: playwright.Float(60000),
		})

	Row, ErrLocator := bb.Page.QuerySelectorAll("select[name=company]>option")
	if ErrLocator != nil {
		return nil, fmt.Errorf("pp.Locator: could not get entries: %v", ErrLocator)
	}
	for _, tr := range Row {
		val, _ := tr.InnerText()
		if !strings.Contains(val, "Выберите марку автомобиля") {
			colors = append(colors, val)
		}
	}

	return colors, nil
}

// Авторизация на страницу
func (bb *Brauzer) Authorization(Login, Password string) error {
	if _, ErrGoto := bb.Page.Goto("https://infocolor.ru/login/?login=yes", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateCommit,
		Timeout:   playwright.Float(3000),
	}); ErrGoto != nil {
		fmt.Println(fmt.Errorf("page.Goto: %v", ErrGoto))
		// return fmt.Errorf("page.Goto: %v", ErrGoto)
	}
	bb.Page.WaitForSelector(`div[class="bx-auth-note"]`)

	ErrFillLogin := bb.Page.Fill(`input[name=USER_LOGIN]`, Login)
	if ErrFillLogin != nil {
		return fmt.Errorf("pp.Fill: Login: %v", ErrFillLogin)
	}

	ErrFillPass := bb.Page.Fill(`input[name=USER_PASSWORD]`, Password)
	if ErrFillPass != nil {
		return fmt.Errorf("pp.Fill: Password: %v", ErrFillPass)
	}

	ErrCheck := bb.Page.Check(`input[name=USER_REMEMBER]`)
	if ErrCheck != nil {
		return fmt.Errorf("pp.Check: %v", ErrCheck)
	}

	// Нажимаем на кнопку
	ErrClick := bb.Page.Click(`td[class=authorize-submit-cell]>input[value=Войти]`)
	if ErrClick != nil {
		return fmt.Errorf("pp.Click: %v", ErrClick)
	}

	bb.Page.WaitForSelector("p[class=notetext]")

	return nil
}
