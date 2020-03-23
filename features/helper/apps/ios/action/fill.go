package iosaction

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios/driver"
)

/*FillElementByXpath element visible by Xpath selector*/
func FillElementByXpath(locator string, text string) error {
	element := ios.Device.FindByXPath(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByButton element visible by button*/
func FillElementByButton(locator string, text string) error {
	element := ios.Device.FindByButton(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByClass element visible by class*/
func FillElementByClass(locator string, text string) error {
	element := ios.Device.FindByClass(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByID element visible by ID*/
func FillElementByID(locator string, text string) error {
	element := ios.Device.FindByID(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByLabel element visible by label*/
func FillElementByLabel(locator string, text string) error {
	element := ios.Device.FindByLabel(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByLink element visible by link*/
func FillElementByLink(locator string, text string) error {
	element := ios.Device.FindByLink(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByName element visible by class name*/
func FillElementByName(locator string, text string) error {
	element := ios.Device.FindByName(locator).Fill(text)
	helper.LogPanicln(element)

	return nil
}

/*FillElementByText element visible by Xpath selector*/
func FillElementByText(locator string, text string) error {
	element := ios.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
	helper.LogPanicln(element)

	return nil
}
