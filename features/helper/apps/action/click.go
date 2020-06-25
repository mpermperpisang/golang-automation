package action

import (
	"github.com/golang-automation/features/helper"
)

/*ClickByXPath : click element by Xpath selector*/
func (s *Page) ClickByXPath(locator string) error {
	element := s.Action.Device.AllByXPath(locator).Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByText : click element by text in xpath*/
func (s *Page) ClickByText(locator string) error {
	element := s.Action.Device.AllByXPath("//*[contains(@text, '" + locator + "')]").Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByButton : click element by button*/
func (s *Page) ClickByButton(locator string) error {
	element := s.Action.Device.AllByButton(locator).Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByClass : click element by class*/
func (s *Page) ClickByClass(locator string) error {
	element := s.Action.Device.AllByClass(locator).Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByID : click element by class ID*/
func (s *Page) ClickByID(locator string) error {
	element := s.Action.Device.AllByID(locator).Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByLabel : click element by label*/
func (s *Page) ClickByLabel(locator string) error {
	element := s.Action.Device.AllByLabel(locator).Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByLink : click element by link*/
func (s *Page) ClickByLink(locator string) error {
	element := s.Action.Device.AllByLink(locator).Click()
	helper.LogPanicln(element)

	return element
}

/*ClickByName : click element by class name*/
func (s *Page) ClickByName(locator string) error {
	element := s.Action.Device.AllByName(locator).Click()
	helper.LogPanicln(element)

	return element
}
