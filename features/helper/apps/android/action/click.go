package androidaction

import (
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android/driver"
)

/*ClickByXPath click element by Xpath selector*/
func ClickByXPath(locator string) error {
	element := android.Device.AllByXPath(locator).Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByText click element by text in xpath*/
func ClickByText(locator string) error {
	element := android.Device.AllByXPath("//*[contains(@text, '" + locator + "')]").Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByButton click element by button*/
func ClickByButton(locator string) error {
	element := android.Device.AllByButton(locator).Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByClass click element by class*/
func ClickByClass(locator string) error {
	element := android.Device.AllByClass(locator).Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByID click element by class ID*/
func ClickByID(locator string) error {
	element := android.Device.AllByID(locator).Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByLabel click element by label*/
func ClickByLabel(locator string) error {
	element := android.Device.AllByLabel(locator).Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByLink click element by link*/
func ClickByLink(locator string) error {
	element := android.Device.AllByLink(locator).Click()
	helper.LogPanicln(element)

	return nil
}

/*ClickByName click element by class name*/
func ClickByName(locator string) error {
	element := android.Device.AllByName(locator).Click()
	helper.LogPanicln(element)

	return nil
}
