package iosaction

import (
	"github.com/golang-automation/features/helper"
)

/*FillElementByXpath element visible by Xpath selector*/
func (s *Page) FillElementByXpath(locator string, text string) error {
	element := s.Action.Device.FindByXPath(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByButton element visible by button*/
func (s *Page) FillElementByButton(locator string, text string) error {
	element := s.Action.Device.FindByButton(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByClass element visible by class*/
func (s *Page) FillElementByClass(locator string, text string) error {
	element := s.Action.Device.FindByClass(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByID element visible by ID*/
func (s *Page) FillElementByID(locator string, text string) error {
	element := s.Action.Device.FindByID(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByLabel element visible by label*/
func (s *Page) FillElementByLabel(locator string, text string) error {
	element := s.Action.Device.FindByLabel(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByLink element visible by link*/
func (s *Page) FillElementByLink(locator string, text string) error {
	element := s.Action.Device.FindByLink(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByName element visible by class name*/
func (s *Page) FillElementByName(locator string, text string) error {
	element := s.Action.Device.FindByName(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

/*FillElementByText element visible by Xpath selector*/
func (s *Page) FillElementByText(locator string, text string) error {
	element := s.Action.Device.FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
	helper.LogPanicln(element)

	return element
}
