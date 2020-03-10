package androidaction

import (
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android"
)

/*FillElementByXpath element visible by Xpath selector*/
func FillElementByXpath(locator string, text string) error {
	err := android.Device.FindByXPath(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByButton element visible by button*/
func FillElementByButton(locator string, text string) error {
	err := android.Device.FindByButton(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByClass element visible by class*/
func FillElementByClass(locator string, text string) error {
	err := android.Device.FindByClass(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByID element visible by ID*/
func FillElementByID(locator string, text string) error {
	err := android.Device.FindByID(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByLabel element visible by label*/
func FillElementByLabel(locator string, text string) error {
	err := android.Device.FindByLabel(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByLink element visible by link*/
func FillElementByLink(locator string, text string) error {
	err := android.Device.FindByLink(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByName element visible by class name*/
func FillElementByName(locator string, text string) error {
	err := android.Device.FindByName(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByText element visible by Xpath selector*/
func FillElementByText(locator string, text string) error {
	err := android.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
	helper.LogPanicln(err)

	return nil
}
