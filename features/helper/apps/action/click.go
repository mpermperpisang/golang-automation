package appsaction

import "github.com/golang-automation/features/helper/errors"

// ClickByXPath : click element by Xpath selector
func (s *Page) ClickByXPath(locator string) error {
	element := s.device().AllByXPath(locator).Click()
	errors.LogPanicln(element)

	return element
}

// ClickByText : click element by text in xpath
func (s *Page) ClickByText(locator string) error {
	element := s.device().AllByXPath("//*[contains(@text, '" + locator + "')]").Click()
	errors.LogPanicln(element)

	return element
}

// ClickByButton : click element by button
func (s *Page) ClickByButton(locator string) error {
	element := s.device().AllByButton(locator).Click()
	errors.LogPanicln(element)

	return element
}

// ClickByClass : click element by class
func (s *Page) ClickByClass(locator string) error {
	element := s.device().AllByClass(locator).Click()
	errors.LogPanicln(element)

	return element
}

// ClickByID : click element by class ID
func (s *Page) ClickByID(locator string) error {
	element := s.device().AllByID(locator).Click()
	errors.LogPanicln(element)

	return element
}

// ClickByLabel : click element by label
func (s *Page) ClickByLabel(locator string) error {
	element := s.device().AllByLabel(locator).Click()
	errors.LogPanicln(element)

	return element
}

// ClickByLink : click element by link
func (s *Page) ClickByLink(locator string) error {
	element := s.device().AllByLink(locator).Click()
	errors.LogPanicln(element)

	return element
}

// ClickByName : click element by class name
func (s *Page) ClickByName(locator string) error {
	element := s.device().AllByName(locator).Click()
	errors.LogPanicln(element)

	return element
}
