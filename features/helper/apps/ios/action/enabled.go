package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementEnabledByXpath element visible by Xpath selector*/
func (s *Page) IsElementEnabledByXpath(locator string) error {
	_, err := s.Action.Device.FindByXPath(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByButton element visible by button*/
func (s *Page) IsElementEnabledByButton(locator string) error {
	_, err := s.Action.Device.FindByButton(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByClass element visible by class*/
func (s *Page) IsElementEnabledByClass(locator string) error {
	_, err := s.Action.Device.FindByClass(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByID element visible by ID*/
func (s *Page) IsElementEnabledByID(locator string) error {
	_, err := s.Action.Device.FindByID(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByLabel element visible by label*/
func (s *Page) IsElementEnabledByLabel(locator string) error {
	_, err := s.Action.Device.FindByLabel(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByLink element visible by link*/
func (s *Page) IsElementEnabledByLink(locator string) error {
	_, err := s.Action.Device.FindByLink(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByName element visible by class name*/
func (s *Page) IsElementEnabledByName(locator string) error {
	_, err := s.Action.Device.FindByName(locator).Enabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByText element visible by Xpath selector*/
func (s *Page) IsElementEnabledByText(locator string) error {
	_, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Enabled()
	helper.LogPanicln(err)

	return nil
}
