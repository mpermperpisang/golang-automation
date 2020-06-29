package appsaction

import "github.com/golang-automation/features/helper/errors"

// PressBack : back to previous page
func (s *Page) PressBack() error {
	element := s.device().Back()
	errors.LogPanicln(element)

	return element
}

// ClearCookies : clear application cookie
func (s *Page) ClearCookies() error {
	element := s.device().ClearCookies()
	errors.LogPanicln(element)

	return element
}

// CancelPopup : cancel popup window
func (s *Page) CancelPopup() error {
	element := s.device().CancelPopup()
	errors.LogPanicln(element)

	return element
}

// CloseApp : close the ios application
func (s *Page) CloseApp() error {
	element := s.device().CloseApp()
	errors.LogPanicln(element)

	return element
}

// ConfirmPopup : confirm popup window
func (s *Page) ConfirmPopup() error {
	element := s.device().ConfirmPopup()
	errors.LogPanicln(element)

	return element
}

// DoubleClick : click twice
func (s *Page) DoubleClick() error {
	element := s.device().DoubleClick()
	errors.LogPanicln(element)

	return element
}

// InstallApp : install the ios application
func (s *Page) InstallApp(path string) error {
	element := s.device().InstallApp(path)
	errors.LogPanicln(element)

	return element
}

// LaunchApp : launch the ios application
func (s *Page) LaunchApp() error {
	element := s.device().LaunchApp()
	errors.LogPanicln(element)

	return element
}

// URLNavigate : navigate app to some url
func (s *Page) URLNavigate(url string) error {
	element := s.device().Navigate(url)
	errors.LogPanicln(element)

	return element
}

// RefreshApp : refresh the ios application
func (s *Page) RefreshApp() error {
	element := s.device().Refresh()
	errors.LogPanicln(element)

	return element
}

// ResetApp : reset the ios application
func (s *Page) ResetApp() error {
	element := s.device().Reset()
	errors.LogPanicln(element)

	return element
}

// TakeScreenshot : take error screenshot
func (s *Page) TakeScreenshot(filename string) error {
	element := s.device().Screenshot(filename)
	errors.LogPanicln(element)

	return element
}
