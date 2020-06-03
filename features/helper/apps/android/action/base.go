package androidaction

import (
	"github.com/golang-automation/features/helper"
)

/*PressBack : back to previous page*/
func (s *Page) PressBack() error {
	element := s.Action.Device.Back()
	helper.LogPanicln(element)

	return element
}

/*ClearCookies : clear application cookie*/
func (s *Page) ClearCookies() error {
	element := s.Action.Device.ClearCookies()
	helper.LogPanicln(element)

	return element
}

/*CancelPopup : cancel popup window*/
func (s *Page) CancelPopup() error {
	element := s.Action.Device.CancelPopup()
	helper.LogPanicln(element)

	return element
}

/*CloseApp : close the android application*/
func (s *Page) CloseApp() error {
	element := s.Action.Device.CloseApp()
	helper.LogPanicln(element)

	return element
}

/*ConfirmPopup : confirm popup window*/
func (s *Page) ConfirmPopup() error {
	element := s.Action.Device.ConfirmPopup()
	helper.LogPanicln(element)

	return element
}

/*DoubleClick : click twice*/
func (s *Page) DoubleClick() error {
	element := s.Action.Device.DoubleClick()
	helper.LogPanicln(element)

	return element
}

/*InstallApp : install the android application*/
func (s *Page) InstallApp(path string) error {
	element := s.Action.Device.InstallApp(path)
	helper.LogPanicln(element)

	return element
}

/*LaunchApp : launch the android application*/
func (s *Page) LaunchApp() error {
	element := s.Action.Device.LaunchApp()
	helper.LogPanicln(element)

	return element
}

/*URLNavigate : navigate app to some url*/
func (s *Page) URLNavigate(url string) error {
	element := s.Action.Device.Navigate(url)
	helper.LogPanicln(element)

	return element
}

/*RefreshApp : refresh the android application*/
func (s *Page) RefreshApp() error {
	element := s.Action.Device.Refresh()
	helper.LogPanicln(element)

	return element
}

/*ResetApp : reset the android application*/
func (s *Page) ResetApp() error {
	element := s.Action.Device.Reset()
	helper.LogPanicln(element)

	return element
}

/*TakeScreenshot : take error screenshot*/
func (s *Page) TakeScreenshot(filename string) error {
	element := s.Action.Device.Screenshot(filename)
	helper.LogPanicln(element)

	return element
}
