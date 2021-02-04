package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// IsElementSelectedByCSS : element selected by CSS selector
func (s Page) IsElementSelectedByCSS(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByID : element selected by class ID
func (s Page) IsElementSelectedByID(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByXpath : element selected by Xpath selector
func (s Page) IsElementSelectedByXpath(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByLinkText : element selected by link text
func (s Page) IsElementSelectedByLinkText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByPartialLink : element selected by partial link text
func (s Page) IsElementSelectedByPartialLink(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByName : element selected by name of class
func (s Page) IsElementSelectedByName(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByTag : element selected by name tag
func (s Page) IsElementSelectedByTag(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByClass : element selected by class name
func (s Page) IsElementSelectedByClass(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

// IsElementSelectedByText : element selected by text in xpath
func (s Page) IsElementSelectedByText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(@text, '"+locator+"')]")
	helper.LogPanicln(err)

	return element.IsSelected()
}
