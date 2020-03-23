package androidaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementVisibleByXpath element visible by Xpath selector*/
func IsElementVisibleByXpath(locator string) error {
	_, err := Device.FindByXPath(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByButton element visible by button*/
func IsElementVisibleByButton(locator string) error {
	_, err := Device.FindByButton(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByClass element visible by class*/
func IsElementVisibleByClass(locator string) error {
	_, err := Device.FindByClass(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByID element visible by ID*/
func IsElementVisibleByID(locator string) error {
	_, err := Device.FindByID(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByLabel element visible by label*/
func IsElementVisibleByLabel(locator string) error {
	_, err := Device.FindByLabel(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByLink element visible by link*/
func IsElementVisibleByLink(locator string) error {
	_, err := Device.FindByLink(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByName element visible by class name*/
func IsElementVisibleByName(locator string) error {
	_, err := Device.FindByName(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByText element visible by Xpath selector*/
func IsElementVisibleByText(locator string) error {
	_, err := Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Visible()
	helper.LogPanicln(err)

	return nil
}
