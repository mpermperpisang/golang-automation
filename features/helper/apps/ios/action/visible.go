package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementVisibleByXpath element visible by Xpath selector*/
func (s *Page) IsElementVisibleByXpath(locator string) bool {
	element, err := s.Action.Device.FindByXPath(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByButton element visible by button*/
func (s *Page) IsElementVisibleByButton(locator string) bool {
	element, err := s.Action.Device.FindByButton(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByClass element visible by class*/
func (s *Page) IsElementVisibleByClass(locator string) bool {
	element, err := s.Action.Device.FindByClass(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByID element visible by ID*/
func (s *Page) IsElementVisibleByID(locator string) bool {
	element, err := s.Action.Device.FindByID(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByLabel element visible by label*/
func (s *Page) IsElementVisibleByLabel(locator string) bool {
	element, err := s.Action.Device.FindByLabel(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByLink element visible by link*/
func (s *Page) IsElementVisibleByLink(locator string) bool {
	element, err := s.Action.Device.FindByLink(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByName element visible by class name*/
func (s *Page) IsElementVisibleByName(locator string) bool {
	element, err := s.Action.Device.FindByName(locator).Visible()
	helper.LogPanicln(err)

	return element
}

/*IsElementVisibleByText element visible by Xpath selector*/
func (s *Page) IsElementVisibleByText(locator string) bool {
	element, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Visible()
	helper.LogPanicln(err)

	return element
}
