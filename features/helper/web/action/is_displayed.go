package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// IsElementDisplayedByCSS : element displayed by CSS selector
func (s *Page) IsElementDisplayedByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByID : element displayed by class ID
func (s *Page) IsElementDisplayedByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByXpath : element displayed by Xpath selector
func (s *Page) IsElementDisplayedByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByLinkText : element displayed by link text
func (s *Page) IsElementDisplayedByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByPartialLink : element displayed by partial link text
func (s *Page) IsElementDisplayedByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByName : element displayed by name of class
func (s *Page) IsElementDisplayedByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByTag : element displayed by name tag
func (s *Page) IsElementDisplayedByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByClass : element displayed by class name
func (s *Page) IsElementDisplayedByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByText : element displayed by text in xpath
func (s *Page) IsElementDisplayedByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}
