package iosaction

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios"
)

/*PressBack : back to previous page*/
func PressBack() error {
	element := ios.Device.Back()
	helper.LogPanicln(element)

	return nil
}

/*ClearCookies : clear application cookie*/
func ClearCookies() error {
	element := ios.Device.ClearCookies()
	helper.LogPanicln(element)

	return nil
}

/*CancelPopup : cancel popup window*/
func CancelPopup() error {
	element := ios.Device.CancelPopup()
	helper.LogPanicln(element)

	return nil
}

/*CloseApp : close the ios application*/
func CloseApp() error {
	element := ios.Device.CloseApp()
	helper.LogPanicln(element)

	return nil
}

/*ConfirmPopup : confirm popup window*/
func ConfirmPopup() error {
	element := ios.Device.ConfirmPopup()
	helper.LogPanicln(element)

	return nil
}

/*DoubleClick : click twice*/
func DoubleClick() error {
	element := ios.Device.DoubleClick()
	helper.LogPanicln(element)

	return nil
}

/*InstallApp : install the ios application*/
func InstallApp(path string) error {
	element := ios.Device.InstallApp(path)
	helper.LogPanicln(element)

	return nil
}

/*LaunchApp : launch the ios application*/
func LaunchApp() error {
	element := ios.Device.LaunchApp()
	helper.LogPanicln(element)

	return nil
}

/*URLNavigate : navigate app to some url*/
func URLNavigate(url string) error {
	element := ios.Device.Navigate(url)
	helper.LogPanicln(element)

	return nil
}

/*RefreshApp : refresh the ios application*/
func RefreshApp() error {
	element := ios.Device.Refresh()
	helper.LogPanicln(element)

	return nil
}

/*ResetApp : reset the ios application*/
func ResetApp() error {
	element := ios.Device.Reset()
	helper.LogPanicln(element)

	return nil
}
