package androidaction

import (
	"fmt"
	"log"

	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device
var err error

/*StartDriver : start android driver*/
func StartDriver() error {
	if err := android.Driver.Start(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*NewDevice : create new android device*/
func NewDevice() error {
	if Device, err = android.Driver.NewDevice(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*PressBack : back to previous page*/
func PressBack() error {
	element := Device.Back()
	helper.LogPanicln(element)

	return nil
}

/*ClearCookies : clear application cookie*/
func ClearCookies() error {
	element := Device.ClearCookies()
	helper.LogPanicln(element)

	return nil
}

/*CancelPopup : cancel popup window*/
func CancelPopup() error {
	element := Device.CancelPopup()
	helper.LogPanicln(element)

	return nil
}

/*CloseApp : close the android application*/
func CloseApp() error {
	element := Device.CloseApp()
	helper.LogPanicln(element)

	return nil
}

/*ConfirmPopup : confirm popup window*/
func ConfirmPopup() error {
	element := Device.ConfirmPopup()
	helper.LogPanicln(element)

	return nil
}

/*DoubleClick : click twice*/
func DoubleClick() error {
	element := Device.DoubleClick()
	helper.LogPanicln(element)

	return nil
}

/*InstallApp : install the android application*/
func InstallApp(path string) error {
	element := Device.InstallApp(path)
	helper.LogPanicln(element)

	return nil
}

/*LaunchApp : launch the android application*/
func LaunchApp() error {
	element := Device.LaunchApp()
	helper.LogPanicln(element)

	return nil
}

/*URLNavigate : navigate app to some url*/
func URLNavigate(url string) error {
	element := Device.Navigate(url)
	helper.LogPanicln(element)

	return nil
}

/*RefreshApp : refresh the android application*/
func RefreshApp() error {
	element := Device.Refresh()
	helper.LogPanicln(element)

	return nil
}

/*ResetApp : reset the android application*/
func ResetApp() error {
	element := Device.Reset()
	helper.LogPanicln(element)

	return nil
}
