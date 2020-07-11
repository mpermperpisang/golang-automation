package appsaction

import "github.com/golang-automation/features/helper"

// PressBack : back to previous page
func (s *Page) PressBack() error {
	element := s.device().Back()
	helper.LogPanicln(element)

	return element
}

// ClearCookies : clear application cookie
func (s *Page) ClearCookies() error {
	element := s.device().ClearCookies()
	helper.LogPanicln(element)

	return element
}

// CancelPopup : cancel popup window
func (s *Page) CancelPopup() error {
	element := s.device().CancelPopup()
	helper.LogPanicln(element)

	return element
}

// CloseApp : close the ios application
func (s *Page) CloseApp() error {
	element := s.device().CloseApp()
	helper.LogPanicln(element)

	return element
}

// ConfirmPopup : confirm popup window
func (s *Page) ConfirmPopup() error {
	element := s.device().ConfirmPopup()
	helper.LogPanicln(element)

	return element
}

// DoubleClick : click twice
func (s *Page) DoubleClick() error {
	element := s.device().DoubleClick()
	helper.LogPanicln(element)

	return element
}

// InstallApp : install the ios application
func (s *Page) InstallApp(path string) error {
	element := s.device().InstallApp(path)
	helper.LogPanicln(element)

	return element
}

// LaunchApp : launch the ios application
func (s *Page) LaunchApp() error {
	element := s.device().LaunchApp()
	helper.LogPanicln(element)

	return element
}

// URLNavigate : navigate app to some url
func (s *Page) URLNavigate(url string) error {
	element := s.device().Navigate(url)
	helper.LogPanicln(element)

	return element
}

// RefreshApp : refresh the ios application
func (s *Page) RefreshApp() error {
	element := s.device().Refresh()
	helper.LogPanicln(element)

	return element
}

// ResetApp : reset the ios application
func (s *Page) ResetApp() error {
	element := s.device().Reset()
	helper.LogPanicln(element)

	return element
}

// TakeScreenshot : take error screenshot
func (s *Page) TakeScreenshot(filename string) error {
	element := s.device().Screenshot(filename)
	helper.LogPanicln(element)

	return element
}
