package appsaction

import "github.com/golang-automation/features/helper"

// FillElementByXpath : fill element by Xpath selector
func (s *Page) FillElementByXpath(locator string, text string) error {
	element := s.device().FindByXPath(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByButton : fill element by button
func (s *Page) FillElementByButton(locator string, text string) error {
	element := s.device().FindByButton(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByClass : fill element by class
func (s *Page) FillElementByClass(locator string, text string) error {
	element := s.device().FindByClass(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByID : fill element by ID
func (s *Page) FillElementByID(locator string, text string) error {
	element := s.device().FindByID(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByLabel : fill element by label
func (s *Page) FillElementByLabel(locator string, text string) error {
	element := s.device().FindByLabel(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByLink : fill element by link
func (s *Page) FillElementByLink(locator string, text string) error {
	element := s.device().FindByLink(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByName : fill element by class name
func (s *Page) FillElementByName(locator string, text string) error {
	element := s.device().FindByName(locator).Fill(text)
	helper.LogPanicln(element)

	return element
}

// FillElementByText : fill element by Xpath selector
func (s *Page) FillElementByText(locator string, text string) error {
	element := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
	helper.LogPanicln(element)

	return element
}
