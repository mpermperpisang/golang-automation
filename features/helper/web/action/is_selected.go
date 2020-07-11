package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// IsElementSelectedByCSS : element selected by CSS selector
func (s *Page) IsElementSelectedByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByID : element selected by class ID
func (s *Page) IsElementSelectedByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByXpath : element selected by Xpath selector
func (s *Page) IsElementSelectedByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByLinkText : element selected by link text
func (s *Page) IsElementSelectedByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByPartialLink : element selected by partial link text
func (s *Page) IsElementSelectedByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByName : element selected by name of class
func (s *Page) IsElementSelectedByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByTag : element selected by name tag
func (s *Page) IsElementSelectedByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByClass : element selected by class name
func (s *Page) IsElementSelectedByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}

// IsElementSelectedByText : element selected by text in xpath
func (s *Page) IsElementSelectedByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.IsSelected()

	return element
}
