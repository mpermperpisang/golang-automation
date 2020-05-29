package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementSelectedByXpath element visible by Xpath selector*/
func (s *Page) IsElementSelectedByXpath(locator string) error {
	_, err := s.Action.Device.FindByXPath(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByButton element visible by button*/
func (s *Page) IsElementSelectedByButton(locator string) error {
	_, err := s.Action.Device.FindByButton(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByClass element visible by class*/
func (s *Page) IsElementSelectedByClass(locator string) error {
	_, err := s.Action.Device.FindByClass(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByID element visible by ID*/
func (s *Page) IsElementSelectedByID(locator string) error {
	_, err := s.Action.Device.FindByID(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByLabel element visible by label*/
func (s *Page) IsElementSelectedByLabel(locator string) error {
	_, err := s.Action.Device.FindByLabel(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByLink element visible by link*/
func (s *Page) IsElementSelectedByLink(locator string) error {
	_, err := s.Action.Device.FindByLink(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByName element visible by class name*/
func (s *Page) IsElementSelectedByName(locator string) error {
	_, err := s.Action.Device.FindByName(locator).Selected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByText element visible by Xpath selector*/
func (s *Page) IsElementSelectedByText(locator string) error {
	_, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Selected()
	helper.LogPanicln(err)

	return nil
}
