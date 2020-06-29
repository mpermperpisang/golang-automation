package webaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// ClearByCSS : clear element by CSS selector
func (s *Page) ClearByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByID : clear element by class ID
func (s *Page) ClearByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByXpath : clear element by Xpath selector
func (s *Page) ClearByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByLinkText : clear element by link text
func (s *Page) ClearByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByPartialLink : clear element by partial link text
func (s *Page) ClearByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByName : clear element by class name
func (s *Page) ClearByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByTag : clear element by name tag
func (s *Page) ClearByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByClass : clear element by class name
func (s *Page) ClearByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	element.Clear()

	return element
}

// ClearByText : clear element by text in xpath
func (s *Page) ClearByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	element.Clear()

	return element
}
