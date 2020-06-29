package webaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// GetAttributeByCSS : get element attribute  by CSS selector
func (s *Page) GetAttributeByCSS(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByID : get element attribute  by class ID
func (s *Page) GetAttributeByID(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByXpath : get element attribute  by Xpath selector
func (s *Page) GetAttributeByXpath(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByLinkText : get element attribute  by link text
func (s *Page) GetAttributeByLinkText(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByPartialLink : get element attribute  by partial link text
func (s *Page) GetAttributeByPartialLink(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByName : get element attribute  by name of class
func (s *Page) GetAttributeByName(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByTag : get element attribute  by name tag
func (s *Page) GetAttributeByTag(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByClass : get element attribute  by class name
func (s *Page) GetAttributeByClass(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}

// GetAttributeByText : get element attribute  by text in xpath
func (s *Page) GetAttributeByText(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "/// [contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	element.GetAttribute(attr)

	return element
}
