package appsaction

// PressBack : back to previous page
func (s *Page) PressBack() error {
	return s.device().Back()
}

// ClearCookies : clear application cookie
func (s *Page) ClearCookies() error {
	return s.device().ClearCookies()
}

// CancelPopup : cancel popup window
func (s *Page) CancelPopup() error {
	return s.device().CancelPopup()
}

// CloseApp : close the ios application
func (s *Page) CloseApp() error {
	return s.device().CloseApp()
}

// ConfirmPopup : confirm popup window
func (s *Page) ConfirmPopup() error {
	return s.device().ConfirmPopup()
}

// DoubleClick : click twice
func (s *Page) DoubleClick() error {
	return s.device().DoubleClick()
}

// InstallApp : install the ios application
func (s *Page) InstallApp(path string) error {
	return s.device().InstallApp(path)
}

// LaunchApp : launch the ios application
func (s *Page) LaunchApp() error {
	return s.device().LaunchApp()
}

// URLNavigate : navigate app to some url
func (s *Page) URLNavigate(url string) error {
	return s.device().Navigate(url)
}

// RefreshApp : refresh the ios application
func (s *Page) RefreshApp() error {
	return s.device().Refresh()
}

// ResetApp : reset the ios application
func (s *Page) ResetApp() error {
	return s.device().Reset()
}

// TakeScreenshot : take error screenshot
func (s *Page) TakeScreenshot(filename string) error {
	return s.device().Screenshot(filename)
}
