package goinfocolor

import (
	"fmt"

	playwright "github.com/playwright-community/playwright-go"
)

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
