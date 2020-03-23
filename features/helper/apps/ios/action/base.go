package iosaction

import (
	"fmt"
	"log"

	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device
var err error

/*StartDriver : start ios driver*/
func StartDriver() error {
	if err := ios.Driver.Start(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*NewDevice : create new ios device*/
func NewDevice() error {
	if Device, err = ios.Driver.NewDevice(); err != nil {
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

/*CloseApp : close the ios application*/
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

/*InstallApp : install the ios application*/
func InstallApp(path string) error {
	element := Device.InstallApp(path)
	helper.LogPanicln(element)

	return nil
}

/*LaunchApp : launch the ios application*/
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

/*RefreshApp : refresh the ios application*/
func RefreshApp() error {
	element := Device.Refresh()
	helper.LogPanicln(element)

	return nil
}

/*ResetApp : reset the ios application*/
func ResetApp() error {
	element := Device.Reset()
	helper.LogPanicln(element)

	return nil
}
