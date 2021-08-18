package actions

var (
	element bool
	err     error
)

func (s Page) PressBack() error {
	return s.device().Back()
}

func (s Page) ClearCookies() error {
	return s.device().ClearCookies()
}

func (s Page) CancelPopup() error {
	return s.device().CancelPopup()
}

func (s Page) CloseApp() error {
	return s.device().CloseApp()
}

func (s Page) ConfirmPopup() error {
	return s.device().ConfirmPopup()
}

func (s Page) DoubleClick() error {
	return s.device().DoubleClick()
}

func (s Page) InstallApp(path string) error {
	return s.device().InstallApp(path)
}

func (s Page) LaunchApp() error {
	return s.device().LaunchApp()
}

func (s Page) URLNavigate(url string) error {
	return s.device().Navigate(url)
}

func (s Page) RefreshApp() error {
	return s.device().Refresh()
}

func (s Page) ResetApp() error {
	return s.device().Reset()
}

func (s Page) TakeScreenshot(filename string) error {
	return s.device().Screenshot(filename)
}
