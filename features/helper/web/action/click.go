package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*ClickByCSS click element by CSS selector*/
func (s *Page) ClickByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByID click element by class ID*/
func (s *Page) ClickByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByXpath click element by Xpath selector*/
func (s *Page) ClickByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByLinkText click element by link text*/
func (s *Page) ClickByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByPartialLink click element by partial link text*/
func (s *Page) ClickByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByName click element by class name*/
func (s *Page) ClickByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByTag click element by name tag*/
func (s *Page) ClickByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByClass click element by class name*/
func (s *Page) ClickByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.Click()

	return element
}

/*ClickByText click element by text in xpath*/
func (s *Page) ClickByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.Click()

	return element
}
