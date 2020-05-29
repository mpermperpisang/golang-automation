package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*GetAttributeByXpath element visible by Xpath selector*/
func (s *Page) GetAttributeByXpath(locator string, attr string) error {
	_, err := s.Action.Device.FindByXPath(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByButton element visible by button*/
func (s *Page) GetAttributeByButton(locator string, attr string) error {
	_, err := s.Action.Device.FindByButton(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByClass element visible by class*/
func (s *Page) GetAttributeByClass(locator string, attr string) error {
	_, err := s.Action.Device.FindByClass(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByID element visible by ID*/
func (s *Page) GetAttributeByID(locator string, attr string) error {
	_, err := s.Action.Device.FindByID(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLabel element visible by label*/
func (s *Page) GetAttributeByLabel(locator string, attr string) error {
	_, err := s.Action.Device.FindByLabel(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLink element visible by link*/
func (s *Page) GetAttributeByLink(locator string, attr string) error {
	_, err := s.Action.Device.FindByLink(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByName element visible by class name*/
func (s *Page) GetAttributeByName(locator string, attr string) error {
	_, err := s.Action.Device.FindByName(locator).Attribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByText element visible by Xpath selector*/
func (s *Page) GetAttributeByText(locator string, attr string) error {
	_, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Attribute(attr)
	helper.LogPanicln(err)

	return nil
}
