package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementVisibleByXpath element visible by Xpath selector*/
func (s *Page) IsElementVisibleByXpath(locator string) error {
	_, err := s.Action.Device.FindByXPath(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByButton element visible by button*/
func (s *Page) IsElementVisibleByButton(locator string) error {
	_, err := s.Action.Device.FindByButton(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByClass element visible by class*/
func (s *Page) IsElementVisibleByClass(locator string) error {
	_, err := s.Action.Device.FindByClass(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByID element visible by ID*/
func (s *Page) IsElementVisibleByID(locator string) error {
	_, err := s.Action.Device.FindByID(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByLabel element visible by label*/
func (s *Page) IsElementVisibleByLabel(locator string) error {
	_, err := s.Action.Device.FindByLabel(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByLink element visible by link*/
func (s *Page) IsElementVisibleByLink(locator string) error {
	_, err := s.Action.Device.FindByLink(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByName element visible by class name*/
func (s *Page) IsElementVisibleByName(locator string) error {
	_, err := s.Action.Device.FindByName(locator).Visible()
	helper.LogPanicln(err)

	return nil
}

/*IsElementVisibleByText element visible by Xpath selector*/
func (s *Page) IsElementVisibleByText(locator string) error {
	_, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Visible()
	helper.LogPanicln(err)

	return nil
}
