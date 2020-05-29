package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementSelectedByXpath element visible by Xpath selector*/
func (s *Page) IsElementSelectedByXpath(locator string) bool {
	element, err := s.Action.Device.FindByXPath(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByButton element visible by button*/
func (s *Page) IsElementSelectedByButton(locator string) bool {
	element, err := s.Action.Device.FindByButton(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByClass element visible by class*/
func (s *Page) IsElementSelectedByClass(locator string) bool {
	element, err := s.Action.Device.FindByClass(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByID element visible by ID*/
func (s *Page) IsElementSelectedByID(locator string) bool {
	element, err := s.Action.Device.FindByID(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByLabel element visible by label*/
func (s *Page) IsElementSelectedByLabel(locator string) bool {
	element, err := s.Action.Device.FindByLabel(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByLink element visible by link*/
func (s *Page) IsElementSelectedByLink(locator string) bool {
	element, err := s.Action.Device.FindByLink(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByName element visible by class name*/
func (s *Page) IsElementSelectedByName(locator string) bool {
	element, err := s.Action.Device.FindByName(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByText element visible by Xpath selector*/
func (s *Page) IsElementSelectedByText(locator string) bool {
	element, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Selected()
	helper.LogPanicln(err)

	return element
}
