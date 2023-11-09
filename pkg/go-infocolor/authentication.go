package goinfocolor

import (
	"fmt"

	playwright "github.com/playwright-community/playwright-go"
)

// Авторизация на страницу
func (bb *Brauzer) Authorization(Login, Password string) error {
	if _, ErrGoto := bb.Page.Goto("https://infocolor.ru/login/?login=yes", playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateCommit,
		Timeout:   playwright.Float(10000),
	}); ErrGoto != nil {
		return fmt.Errorf("page.Goto: %v", ErrGoto)
	}

	bb.Page.WaitForSelector(`div[class="bx-auth-note"]`,
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateHidden,
			Timeout: playwright.Float(10000),
		})

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

	bb.Page.WaitForSelector(`p[class=notetext]`,
		playwright.PageWaitForSelectorOptions{
			State:   playwright.WaitForSelectorStateHidden,
			Timeout: playwright.Float(10000),
		})

	return nil
}
