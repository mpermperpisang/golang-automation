package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// IsElementDisplayedByCSS : element displayed by CSS selector
func (s Page) IsElementDisplayedByCSS(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByID : element displayed by class ID
func (s Page) IsElementDisplayedByID(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByXpath : element displayed by Xpath selector
func (s Page) IsElementDisplayedByXpath(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByLinkText : element displayed by link text
func (s Page) IsElementDisplayedByLinkText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByPartialLink : element displayed by partial link text
func (s Page) IsElementDisplayedByPartialLink(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByName : element displayed by name of class
func (s Page) IsElementDisplayedByName(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByTag : element displayed by name tag
func (s Page) IsElementDisplayedByTag(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByClass : element displayed by class name
func (s Page) IsElementDisplayedByClass(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

// IsElementDisplayedByText : element displayed by text in xpath
func (s Page) IsElementDisplayedByText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(@text, '"+locator+"')]")
	helper.LogPanicln(err)

	return element.IsDisplayed()
}
