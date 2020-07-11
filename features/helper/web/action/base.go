package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// AcceptAlert : accepting alert
func (s *Page) AcceptAlert() error {
	element := s.driver().AcceptAlert()
	helper.LogPanicln(element)

	return element
}

// DismissAlert : dismissing alert
func (s *Page) DismissAlert() error {
	element := s.driver().DismissAlert()
	helper.LogPanicln(element)

	return element
}

// BackBrowser : back to previous page
func (s *Page) BackBrowser() error {
	element := s.driver().Back()
	helper.LogPanicln(element)

	return element
}

// GetCurrentURL : get current browser url
func (s *Page) GetCurrentURL() string {
	element, err := s.driver().CurrentURL()
	helper.LogPanicln(err)

	return element
}

// ButtonDown : button down
func (s *Page) ButtonDown() error {
	element := s.driver().ButtonDown()
	helper.LogPanicln(element)

	return element
}

// ButtonUp : buttom up
func (s *Page) ButtonUp() error {
	element := s.driver().ButtonUp()
	helper.LogPanicln(element)

	return element
}

// CloseWindow : close browser window
func (s *Page) CloseWindow(locator string) error {
	element := s.driver().CloseWindow(locator)
	helper.LogPanicln(element)

	return element
}

// AddCookie : add cookies in browser
func (s *Page) AddCookie(cookie *selenium.Cookie) error {
	element := s.driver().AddCookie(cookie)
	helper.LogPanicln(element)

	return element
}

// DeleteCookie : delete cookies in browser
func (s *Page) DeleteCookie(locator string) error {
	element := s.driver().DeleteCookie(locator)
	helper.LogPanicln(element)

	return element
}

// DeleteAllCookies : delete all cookies in browser
func (s *Page) DeleteAllCookies(locator string) error {
	element := s.driver().DeleteAllCookies()
	helper.LogPanicln(element)

	return element
}

// DoubleClick : double click in browser
func (s *Page) DoubleClick() error {
	element := s.driver().DoubleClick()
	helper.LogPanicln(element)

	return element
}

// Refresh : refresh browser
func (s *Page) Refresh() error {
	element := s.driver().Refresh()
	helper.LogPanicln(element)

	return element
}

// ResizeWindow : set browser size
func (s *Page) ResizeWindow(locator string, w int, h int) error {
	element := s.driver().ResizeWindow(locator, w, h)
	helper.LogPanicln(element)

	return element
}

// MaxWindow : set browser to maximum size
func (s *Page) MaxWindow(locator string) error {
	element := s.driver().MaximizeWindow(locator)
	helper.LogPanicln(element)

	return element
}

// TakeScreenshot : take error screenshot
func (s *Page) TakeScreenshot() []byte {
	element, err := s.driver().Screenshot()
	helper.LogPanicln(err)

	return element
}
