package action

import (
	"github.com/golang-automation/features/helper"
)

/*IsElementSelectedByXpath : element selected by Xpath selector*/
func (s *Page) IsElementSelectedByXpath(locator string) bool {
	element, err := s.Action.Device.FindByXPath(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByButton : element selected by button*/
func (s *Page) IsElementSelectedByButton(locator string) bool {
	element, err := s.Action.Device.FindByButton(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByClass : element selected by class*/
func (s *Page) IsElementSelectedByClass(locator string) bool {
	element, err := s.Action.Device.FindByClass(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByID : element selected by ID*/
func (s *Page) IsElementSelectedByID(locator string) bool {
	element, err := s.Action.Device.FindByID(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByLabel : element selected by label*/
func (s *Page) IsElementSelectedByLabel(locator string) bool {
	element, err := s.Action.Device.FindByLabel(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByLink : element selected by link*/
func (s *Page) IsElementSelectedByLink(locator string) bool {
	element, err := s.Action.Device.FindByLink(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByName : element selected by class name*/
func (s *Page) IsElementSelectedByName(locator string) bool {
	element, err := s.Action.Device.FindByName(locator).Selected()
	helper.LogPanicln(err)

	return element
}

/*IsElementSelectedByText : element selected by Xpath selector*/
func (s *Page) IsElementSelectedByText(locator string) bool {
	element, err := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Selected()
	helper.LogPanicln(err)

	return element
}
