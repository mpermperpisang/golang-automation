package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// GetAttributeByCSS : get element attribute  by CSS selector
func (s Page) GetAttributeByCSS(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByID : get element attribute  by class ID
func (s Page) GetAttributeByID(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByXpath : get element attribute  by Xpath selector
func (s Page) GetAttributeByXpath(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByLinkText : get element attribute  by link text
func (s Page) GetAttributeByLinkText(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByPartialLink : get element attribute  by partial link text
func (s Page) GetAttributeByPartialLink(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByName : get element attribute  by name of class
func (s Page) GetAttributeByName(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByTag : get element attribute  by name tag
func (s Page) GetAttributeByTag(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByClass : get element attribute  by class name
func (s Page) GetAttributeByClass(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

// GetAttributeByText : get element attribute  by text in xpath
func (s Page) GetAttributeByText(locator string, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(@text, '"+locator+"')]")
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}
