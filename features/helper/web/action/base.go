package webaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// AcceptAlert : accepting alert
func (s *Page) AcceptAlert() error {
	element := s.driver().AcceptAlert()
	errors.LogPanicln(element)

	return element
}

// DismissAlert : dismissing alert
func (s *Page) DismissAlert() error {
	element := s.driver().DismissAlert()
	errors.LogPanicln(element)

	return element
}

// BackBrowser : back to previous page
func (s *Page) BackBrowser() error {
	element := s.driver().Back()
	errors.LogPanicln(element)

	return element
}

// GetCurrentURL : get current browser url
func (s *Page) GetCurrentURL() string {
	element, err := s.driver().CurrentURL()
	errors.LogPanicln(err)

	return element
}

// ButtonDown : button down
func (s *Page) ButtonDown() error {
	element := s.driver().ButtonDown()
	errors.LogPanicln(element)

	return element
}

// ButtonUp : buttom up
func (s *Page) ButtonUp() error {
	element := s.driver().ButtonUp()
	errors.LogPanicln(element)

	return element
}

// CloseWindow : close browser window
func (s *Page) CloseWindow(locator string) error {
	element := s.driver().CloseWindow(locator)
	errors.LogPanicln(element)

	return element
}

// AddCookie : add cookies in browser
func (s *Page) AddCookie(cookie *selenium.Cookie) error {
	element := s.driver().AddCookie(cookie)
	errors.LogPanicln(element)

	return element
}

// DeleteCookie : delete cookies in browser
func (s *Page) DeleteCookie(locator string) error {
	element := s.driver().DeleteCookie(locator)
	errors.LogPanicln(element)

	return element
}

// DeleteAllCookies : delete all cookies in browser
func (s *Page) DeleteAllCookies(locator string) error {
	element := s.driver().DeleteAllCookies()
	errors.LogPanicln(element)

	return element
}

// DoubleClick : double click in browser
func (s *Page) DoubleClick() error {
	element := s.driver().DoubleClick()
	errors.LogPanicln(element)

	return element
}

// Refresh : refresh browser
func (s *Page) Refresh() error {
	element := s.driver().Refresh()
	errors.LogPanicln(element)

	return element
}

// ResizeWindow : set browser size
func (s *Page) ResizeWindow(locator string, w int, h int) error {
	element := s.driver().ResizeWindow(locator, w, h)
	errors.LogPanicln(element)

	return element
}

// MaxWindow : set browser to maximum size
func (s *Page) MaxWindow(locator string) error {
	element := s.driver().MaximizeWindow(locator)
	errors.LogPanicln(element)

	return element
}

// TakeScreenshot : take error screenshot
func (s *Page) TakeScreenshot() []byte {
	element, err := s.driver().Screenshot()
	errors.LogPanicln(err)

	return element
}
