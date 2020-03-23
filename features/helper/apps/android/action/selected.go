package androidaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementSelectedByXpath element visible by Xpath selector*/
func IsElementSelectedByXpath(locator string) error {
	_, err := Device.FindByXPath(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByButton element visible by button*/
func IsElementSelectedByButton(locator string) error {
	_, err := Device.FindByButton(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByClass element visible by class*/
func IsElementSelectedByClass(locator string) error {
	_, err := Device.FindByClass(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByID element visible by ID*/
func IsElementSelectedByID(locator string) error {
	_, err := Device.FindByID(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByLabel element visible by label*/
func IsElementSelectedByLabel(locator string) error {
	_, err := Device.FindByLabel(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByLink element visible by link*/
func IsElementSelectedByLink(locator string) error {
	_, err := Device.FindByLink(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByName element visible by class name*/
func IsElementSelectedByName(locator string) error {
	_, err := Device.FindByName(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByText element visible by Xpath selector*/
func IsElementSelectedByText(locator string) error {
	_, err := Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Selected()
	helper.LogPanicln(err)

	return nil
}
