package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*AcceptAlert : accepting alert*/
func (s *Page) AcceptAlert() error {
	element := s.Action.Driver.AcceptAlert()
	helper.LogPanicln(element)

	return element
}

/*DismissAlert : dismissing alert*/
func (s *Page) DismissAlert() error {
	element := s.Action.Driver.DismissAlert()
	helper.LogPanicln(element)

	return element
}

/*BackBrowser : back to previous page*/
func (s *Page) BackBrowser() error {
	element := s.Action.Driver.Back()
	helper.LogPanicln(element)

	return element
}

/*GetCurrentURL : get current browser url*/
func (s *Page) GetCurrentURL() string {
	element, err := s.Action.Driver.CurrentURL()
	helper.LogPanicln(err)

	return element
}

/*ButtonDown : button down*/
func (s *Page) ButtonDown() error {
	element := s.Action.Driver.ButtonDown()
	helper.LogPanicln(element)

	return element
}

/*ButtonUp : buttom up*/
func (s *Page) ButtonUp() error {
	element := s.Action.Driver.ButtonUp()
	helper.LogPanicln(element)

	return element
}

/*CloseWindow : close browser window*/
func (s *Page) CloseWindow(locator string) error {
	element := s.Action.Driver.CloseWindow(locator)
	helper.LogPanicln(element)

	return element
}

/*AddCookie : add cookies in browser*/
func (s *Page) AddCookie(cookie *selenium.Cookie) error {
	element := s.Action.Driver.AddCookie(cookie)
	helper.LogPanicln(element)

	return element
}

/*DeleteCookie : delete cookies in browser*/
func (s *Page) DeleteCookie(locator string) error {
	element := s.Action.Driver.DeleteCookie(locator)
	helper.LogPanicln(element)

	return element
}

/*DeleteAllCookies : delete all cookies in browser*/
func (s *Page) DeleteAllCookies(locator string) error {
	element := s.Action.Driver.DeleteAllCookies()
	helper.LogPanicln(element)

	return element
}

/*DoubleClick : double click in browser*/
func (s *Page) DoubleClick() error {
	element := s.Action.Driver.DoubleClick()
	helper.LogPanicln(element)

	return element
}

/*Refresh : refresh browser*/
func (s *Page) Refresh() error {
	element := s.Action.Driver.Refresh()
	helper.LogPanicln(element)

	return element
}

/*ResizeWindow : set browser size*/
func (s *Page) ResizeWindow(locator string, w int, h int) error {
	element := s.Action.Driver.ResizeWindow(locator, w, h)
	helper.LogPanicln(element)

	return element
}

/*MaxWindow : set browser to maximum size*/
func (s *Page) MaxWindow(locator string) error {
	element := s.Action.Driver.MaximizeWindow(locator)
	helper.LogPanicln(element)

	return element
}
