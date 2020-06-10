package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*ClearByCSS : clear element by CSS selector*/
func (s *Page) ClearByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByID : clear element by class ID*/
func (s *Page) ClearByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByXpath : clear element by Xpath selector*/
func (s *Page) ClearByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByLinkText : clear element by link text*/
func (s *Page) ClearByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByPartialLink : clear element by partial link text*/
func (s *Page) ClearByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByName : clear element by class name*/
func (s *Page) ClearByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByTag : clear element by name tag*/
func (s *Page) ClearByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByClass : clear element by class name*/
func (s *Page) ClearByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.Clear()

	return element
}

/*ClearByText : clear element by text in xpath*/
func (s *Page) ClearByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.Clear()

	return element
}
