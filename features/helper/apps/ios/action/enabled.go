package iosaction

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios/driver"
)

/*IsElementEnabledByXpath element visible by Xpath selector*/
func IsElementEnabledByXpath(locator string) error {
	_, err := ios.Device.FindByXPath(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByButton element visible by button*/
func IsElementEnabledByButton(locator string) error {
	_, err := ios.Device.FindByButton(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByClass element visible by class*/
func IsElementEnabledByClass(locator string) error {
	_, err := ios.Device.FindByClass(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByID element visible by ID*/
func IsElementEnabledByID(locator string) error {
	_, err := ios.Device.FindByID(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByLabel element visible by label*/
func IsElementEnabledByLabel(locator string) error {
	_, err := ios.Device.FindByLabel(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByLink element visible by link*/
func IsElementEnabledByLink(locator string) error {
	_, err := ios.Device.FindByLink(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByName element visible by class name*/
func IsElementEnabledByName(locator string) error {
	_, err := ios.Device.FindByName(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByText element visible by Xpath selector*/
func IsElementEnabledByText(locator string) error {
	_, err := ios.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Enabled()
	helper.LogPanicln(err)

	return nil
}
