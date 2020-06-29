package action

import "github.com/golang-automation/features/helper/errors"

// FillElementByXpath : fill element by Xpath selector
func (s *Page) FillElementByXpath(locator string, text string) error {
	element := s.Action.Device.FindByXPath(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByButton : fill element by button
func (s *Page) FillElementByButton(locator string, text string) error {
	element := s.Action.Device.FindByButton(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByClass : fill element by class
func (s *Page) FillElementByClass(locator string, text string) error {
	element := s.Action.Device.FindByClass(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByID : fill element by ID
func (s *Page) FillElementByID(locator string, text string) error {
	element := s.Action.Device.FindByID(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByLabel : fill element by label
func (s *Page) FillElementByLabel(locator string, text string) error {
	element := s.Action.Device.FindByLabel(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByLink : fill element by link
func (s *Page) FillElementByLink(locator string, text string) error {
	element := s.Action.Device.FindByLink(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByName : fill element by class name
func (s *Page) FillElementByName(locator string, text string) error {
	element := s.Action.Device.FindByName(locator).Fill(text)
	errors.LogPanicln(element)

	return element
}

// FillElementByText : fill element by Xpath selector
func (s *Page) FillElementByText(locator string, text string) error {
	element := s.Action.Device.FindByXPath("/// [contains(@text, '" + locator + "')]").Fill(text)
	errors.LogPanicln(element)

	return element
}
