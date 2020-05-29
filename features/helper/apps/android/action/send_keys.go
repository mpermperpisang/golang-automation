package androidaction

import (
	"github.com/golang-automation/features/helper"
)

/*SendKeysByXpath input element by Xpath selector*/
func (s *Page) SendKeysByXpath(locator string, text string) error {
	element := s.Action.Device.FindByXPath(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByButton input element by button*/
func (s *Page) SendKeysByButton(locator string, text string) error {
	element := s.Action.Device.FindByButton(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByClass input element by class*/
func (s *Page) SendKeysByClass(locator string, text string) error {
	element := s.Action.Device.FindByClass(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByID input element by ID*/
func (s *Page) SendKeysByID(locator string, text string) error {
	element := s.Action.Device.FindByID(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByLabel input element by label*/
func (s *Page) SendKeysByLabel(locator string, text string) error {
	element := s.Action.Device.FindByLabel(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByLink input element by link*/
func (s *Page) SendKeysByLink(locator string, text string) error {
	element := s.Action.Device.FindByLink(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByName input element by class name*/
func (s *Page) SendKeysByName(locator string, text string) error {
	element := s.Action.Device.FindByName(locator).SendKeys(text)
	helper.LogPanicln(element)

	return element
}

/*SendKeysByText input element by text in xpath*/
func (s *Page) SendKeysByText(locator string, text string) error {
	element := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").SendKeys(text)
	helper.LogPanicln(element)

	return element
}
