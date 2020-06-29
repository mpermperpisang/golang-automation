package webaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// IsElementDisplayedByCSS : element displayed by CSS selector
func (s *Page) IsElementDisplayedByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByID : element displayed by class ID
func (s *Page) IsElementDisplayedByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByXpath : element displayed by Xpath selector
func (s *Page) IsElementDisplayedByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByLinkText : element displayed by link text
func (s *Page) IsElementDisplayedByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByPartialLink : element displayed by partial link text
func (s *Page) IsElementDisplayedByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByName : element displayed by name of class
func (s *Page) IsElementDisplayedByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByTag : element displayed by name tag
func (s *Page) IsElementDisplayedByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByClass : element displayed by class name
func (s *Page) IsElementDisplayedByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}

// IsElementDisplayedByText : element displayed by text in xpath
func (s *Page) IsElementDisplayedByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "/// [contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	element.IsDisplayed()

	return element
}
