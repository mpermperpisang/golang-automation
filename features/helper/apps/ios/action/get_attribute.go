package iosaction

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios/driver"
)

/*GetAttributeByXpath element visible by Xpath selector*/
func GetAttributeByXpath(locator string, attr string) error {
	_, err := ios.Device.FindByXPath(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByButton element visible by button*/
func GetAttributeByButton(locator string, attr string) error {
	_, err := ios.Device.FindByButton(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByClass element visible by class*/
func GetAttributeByClass(locator string, attr string) error {
	_, err := ios.Device.FindByClass(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByID element visible by ID*/
func GetAttributeByID(locator string, attr string) error {
	_, err := ios.Device.FindByID(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLabel element visible by label*/
func GetAttributeByLabel(locator string, attr string) error {
	_, err := ios.Device.FindByLabel(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLink element visible by link*/
func GetAttributeByLink(locator string, attr string) error {
	_, err := ios.Device.FindByLink(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByName element visible by class name*/
func GetAttributeByName(locator string, attr string) error {
	_, err := ios.Device.FindByName(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByText element visible by Xpath selector*/
func GetAttributeByText(locator string, attr string) error {
	_, err := ios.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Attribute(attr)
	helper.LogPanicln(err)

	return nil
}
