package appsaction

import "github.com/golang-automation/features/helper"

// GetAttributeByXpath : get element attribute by Xpath selector
func (s *Page) GetAttributeByXpath(locator string, attr string) string {
	element, err := s.device().FindByXPath(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByButton : get element attribute by button
func (s *Page) GetAttributeByButton(locator string, attr string) string {
	element, err := s.device().FindByButton(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByClass : get element attribute by class
func (s *Page) GetAttributeByClass(locator string, attr string) string {
	element, err := s.device().FindByClass(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByID : get element attribute by ID
func (s *Page) GetAttributeByID(locator string, attr string) string {
	element, err := s.device().FindByID(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByLabel : get element attribute by label
func (s *Page) GetAttributeByLabel(locator string, attr string) string {
	element, err := s.device().FindByLabel(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByLink : get element attribute by link
func (s *Page) GetAttributeByLink(locator string, attr string) string {
	element, err := s.device().FindByLink(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByName : get element attribute by class name
func (s *Page) GetAttributeByName(locator string, attr string) string {
	element, err := s.device().FindByName(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

// GetAttributeByText : get element attribute by Xpath selector
func (s *Page) GetAttributeByText(locator string, attr string) string {
	element, err := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Attribute(attr)
	helper.LogPanicln(err)

	return element
}
