package appsaction

import "github.com/golang-automation/features/helper/errors"

// IsElementVisibleByXpath : element visible by Xpath selector
func (s *Page) IsElementVisibleByXpath(locator string) bool {
	element, err := s.device().FindByXPath(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByButton : element visible by button
func (s *Page) IsElementVisibleByButton(locator string) bool {
	element, err := s.device().FindByButton(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByClass : element visible by class
func (s *Page) IsElementVisibleByClass(locator string) bool {
	element, err := s.device().FindByClass(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByID : element visible by ID
func (s *Page) IsElementVisibleByID(locator string) bool {
	element, err := s.device().FindByID(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByLabel : element visible by label
func (s *Page) IsElementVisibleByLabel(locator string) bool {
	element, err := s.device().FindByLabel(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByLink : element visible by link
func (s *Page) IsElementVisibleByLink(locator string) bool {
	element, err := s.device().FindByLink(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByName : element visible by class name
func (s *Page) IsElementVisibleByName(locator string) bool {
	element, err := s.device().FindByName(locator).Visible()
	errors.LogPanicln(err)

	return element
}

// IsElementVisibleByText : element visible by Xpath selector
func (s *Page) IsElementVisibleByText(locator string) bool {
	element, err := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Visible()
	errors.LogPanicln(err)

	return element
}
