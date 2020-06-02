package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementEnabledByXpath : element enabled by Xpath selector*/
func (s *Page) IsElementEnabledByXpath(locator string) bool {
	element, err := s.Action.Device.FindByXPath(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByButton : element enabled by button*/
func (s *Page) IsElementEnabledByButton(locator string) bool {
	element, err := s.Action.Device.FindByButton(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByClass : element enabled by class*/
func (s *Page) IsElementEnabledByClass(locator string) bool {
	element, err := s.Action.Device.FindByClass(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByID : element enabled by ID*/
func (s *Page) IsElementEnabledByID(locator string) bool {
	element, err := s.Action.Device.FindByID(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByLabel : element enabled by label*/
func (s *Page) IsElementEnabledByLabel(locator string) bool {
	element, err := s.Action.Device.FindByLabel(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByLink : element enabled by link*/
func (s *Page) IsElementEnabledByLink(locator string) bool {
	element, err := s.Action.Device.FindByLink(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByName : element enabled by class name*/
func (s *Page) IsElementEnabledByName(locator string) bool {
	element, err := s.Action.Device.FindByName(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

/*IsElementEnabledByText : element enabled by Xpath selector*/
func (s *Page) IsElementEnabledByText(locator string) bool {
	element, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Enabled()
	helper.LogPanicln(err)

	return element
}
