package webaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// GetTextByCSS : get element text by CSS selector
func (s *Page) GetTextByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByID : get element text by class ID
func (s *Page) GetTextByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByXpath : get element text by Xpath selector
func (s *Page) GetTextByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByLinkText : get element text by link text
func (s *Page) GetTextByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByPartialLink : get element text by partial link text
func (s *Page) GetTextByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByName : get element text by name of class
func (s *Page) GetTextByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByText : get element text by text in xpath
func (s *Page) GetTextByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByTag : get element text by name tag
func (s *Page) GetTextByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}

// GetTextByClass : get element text by class name
func (s *Page) GetTextByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	element.Text()

	return element
}
