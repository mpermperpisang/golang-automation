package webaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// IsElementEnabledByCSS : element enabled by CSS selector
func (s *Page) IsElementEnabledByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByID : element enabled by class ID
func (s *Page) IsElementEnabledByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByXpath : element enabled by Xpath selector
func (s *Page) IsElementEnabledByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByLinkText : element enabled by link text
func (s *Page) IsElementEnabledByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByPartialLink : element enabled by partial link text
func (s *Page) IsElementEnabledByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByName : element enabled by name of class
func (s *Page) IsElementEnabledByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByTag : element enabled by name tag
func (s *Page) IsElementEnabledByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByClass : element enabled by class name
func (s *Page) IsElementEnabledByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}

// IsElementEnabledByText : element enabled by text in xpath
func (s *Page) IsElementEnabledByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	element.IsEnabled()

	return element
}
