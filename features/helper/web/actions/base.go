package actions

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

func (s Page) AcceptAlert() error {
	return s.driver().AcceptAlert()
}

func (s Page) DismissAlert() error {
	return s.driver().DismissAlert()
}

func (s Page) BackBrowser() error {
	return s.driver().Back()
}

func (s Page) GetCurrentURL() string {
	element, err := s.driver().CurrentURL()
	helper.LogPanicln(err)

	return element
}

func (s Page) GoToURL(url string) error {
	return s.driver().Get(url)
}

func (s Page) ButtonDown() error {
	return s.driver().ButtonDown()
}

func (s Page) ButtonUp() error {
	return s.driver().ButtonUp()
}

func (s Page) CloseWindow(locator string) error {
	return s.driver().CloseWindow(locator)
}

func (s Page) AddCookie(cookie *selenium.Cookie) error {
	return s.driver().AddCookie(cookie)
}

func (s Page) DeleteCookie(locator string) error {
	return s.driver().DeleteCookie(locator)
}

func (s Page) DeleteAllCookies(locator string) error {
	return s.driver().DeleteAllCookies()
}

func (s Page) DoubleClick() error {
	return s.driver().DoubleClick()
}

func (s Page) Refresh() error {
	return s.driver().Refresh()
}

func (s Page) ResizeWindow(locator string, w, h int) error {
	return s.driver().ResizeWindow(locator, w, h)
}

func (s Page) MaxWindow(locator string) error {
	return s.driver().MaximizeWindow(locator)
}

func (s Page) TakeScreenshot() []byte {
	element, err := s.driver().Screenshot()
	helper.LogPanicln(err)

	return element
}
