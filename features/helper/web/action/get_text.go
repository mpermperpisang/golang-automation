package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// GetTextByCSS : get element text by CSS selector
func (s Page) GetTextByCSS(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByID : get element text by class ID
func (s Page) GetTextByID(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByXpath : get element text by Xpath selector
func (s Page) GetTextByXpath(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByLinkText : get element text by link text
func (s Page) GetTextByLinkText(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByPartialLink : get element text by partial link text
func (s Page) GetTextByPartialLink(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByName : get element text by name of class
func (s Page) GetTextByName(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByText : get element text by text in xpath
func (s Page) GetTextByText(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(@text, '"+locator+"')]")
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByTag : get element text by name tag
func (s Page) GetTextByTag(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.Text()
}

// GetTextByClass : get element text by class name
func (s Page) GetTextByClass(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.Text()
}
