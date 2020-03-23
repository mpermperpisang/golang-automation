package androidaction

import (
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android/driver"
)

/*PressBack : back to previous page*/
func PressBack() error {
	element := android.Device.Back()
	helper.LogPanicln(element)

	return nil
}

/*ClearCookies : clear application cookie*/
func ClearCookies() error {
	element := android.Device.ClearCookies()
	helper.LogPanicln(element)

	return nil
}

/*CancelPopup : cancel popup window*/
func CancelPopup() error {
	element := android.Device.CancelPopup()
	helper.LogPanicln(element)

	return nil
}

/*CloseApp : close the android application*/
func CloseApp() error {
	element := android.Device.CloseApp()
	helper.LogPanicln(element)

	return nil
}

/*ConfirmPopup : confirm popup window*/
func ConfirmPopup() error {
	element := android.Device.ConfirmPopup()
	helper.LogPanicln(element)

	return nil
}

/*DoubleClick : click twice*/
func DoubleClick() error {
	element := android.Device.DoubleClick()
	helper.LogPanicln(element)

	return nil
}

/*InstallApp : install the android application*/
func InstallApp(path string) error {
	element := android.Device.InstallApp(path)
	helper.LogPanicln(element)

	return nil
}

/*LaunchApp : launch the android application*/
func LaunchApp() error {
	element := android.Device.LaunchApp()
	helper.LogPanicln(element)

	return nil
}

/*URLNavigate : navigate app to some url*/
func URLNavigate(url string) error {
	element := android.Device.Navigate(url)
	helper.LogPanicln(element)

	return nil
}

/*RefreshApp : refresh the android application*/
func RefreshApp() error {
	element := android.Device.Refresh()
	helper.LogPanicln(element)

	return nil
}

/*ResetApp : reset the android application*/
func ResetApp() error {
	element := android.Device.Reset()
	helper.LogPanicln(element)

	return nil
}
