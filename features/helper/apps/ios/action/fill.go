package iosaction

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios"
)

/*FillElementByXpath element visible by Xpath selector*/
func FillElementByXpath(locator string, text string) error {
	err := ios.Device.FindByXPath(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByButton element visible by button*/
func FillElementByButton(locator string, text string) error {
	err := ios.Device.FindByButton(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByClass element visible by class*/
func FillElementByClass(locator string, text string) error {
	err := ios.Device.FindByClass(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByID element visible by ID*/
func FillElementByID(locator string, text string) error {
	err := ios.Device.FindByID(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByLabel element visible by label*/
func FillElementByLabel(locator string, text string) error {
	err := ios.Device.FindByLabel(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByLink element visible by link*/
func FillElementByLink(locator string, text string) error {
	err := ios.Device.FindByLink(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByName element visible by class name*/
func FillElementByName(locator string, text string) error {
	err := ios.Device.FindByName(locator).Fill(text)
	helper.LogPanicln(err)

	return nil
}

/*FillElementByText element visible by Xpath selector*/
func FillElementByText(locator string, text string) error {
	err := ios.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
	helper.LogPanicln(err)

	return nil
}
