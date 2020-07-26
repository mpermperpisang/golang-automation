package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// IsElementEnabledByCSS : element enabled by CSS selector
func (s *Page) IsElementEnabledByCSS(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByID : element enabled by class ID
func (s *Page) IsElementEnabledByID(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByXpath : element enabled by Xpath selector
func (s *Page) IsElementEnabledByXpath(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByLinkText : element enabled by link text
func (s *Page) IsElementEnabledByLinkText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByPartialLink : element enabled by partial link text
func (s *Page) IsElementEnabledByPartialLink(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByName : element enabled by name of class
func (s *Page) IsElementEnabledByName(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByTag : element enabled by name tag
func (s *Page) IsElementEnabledByTag(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByClass : element enabled by class name
func (s *Page) IsElementEnabledByClass(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

// IsElementEnabledByText : element enabled by text in xpath
func (s *Page) IsElementEnabledByText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.IsEnabled()
}
