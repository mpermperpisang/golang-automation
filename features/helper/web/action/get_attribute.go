package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*GetAttributeByCSS element enable by CSS selector*/
func (s *Page) GetAttributeByCSS(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByID element enable by class ID*/
func (s *Page) GetAttributeByID(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByXpath element enable by Xpath selector*/
func (s *Page) GetAttributeByXpath(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByLinkText element enable by link text*/
func (s *Page) GetAttributeByLinkText(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByPartialLink element enable by partial link text*/
func (s *Page) GetAttributeByPartialLink(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByName element enable by name of class*/
func (s *Page) GetAttributeByName(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByTag element enable by name tag*/
func (s *Page) GetAttributeByTag(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByClass element enable by class name*/
func (s *Page) GetAttributeByClass(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}

/*GetAttributeByText element enable by text in xpath*/
func (s *Page) GetAttributeByText(locator string, attr string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.GetAttribute(attr)

	return nil
}
