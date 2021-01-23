package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// ClearByCSS : clear element by CSS selector
func (s *Page) ClearByCSS(locator string) error {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByID : clear element by class ID
func (s *Page) ClearByID(locator string) error {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByXpath : clear element by Xpath selector
func (s *Page) ClearByXpath(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByLinkText : clear element by link text
func (s *Page) ClearByLinkText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByPartialLink : clear element by partial link text
func (s *Page) ClearByPartialLink(locator string) error {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByName : clear element by class name
func (s *Page) ClearByName(locator string) error {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByTag : clear element by name tag
func (s *Page) ClearByTag(locator string) error {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByClass : clear element by class name
func (s *Page) ClearByClass(locator string) error {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

// ClearByText : clear element by text in xpath
func (s *Page) ClearByText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(@text, '"+locator+"')]")
	helper.LogPanicln(err)

	return element.Clear()
}
