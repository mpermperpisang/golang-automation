package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*IsElementDisplayedByCSS element display by CSS selector*/
func (s *Page) IsElementDisplayedByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByID element display by class ID*/
func (s *Page) IsElementDisplayedByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByXpath element display by Xpath selector*/
func (s *Page) IsElementDisplayedByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByLinkText element display by link text*/
func (s *Page) IsElementDisplayedByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByPartialLink element display by partial link text*/
func (s *Page) IsElementDisplayedByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByName element display by name of class*/
func (s *Page) IsElementDisplayedByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByTag element display by name tag*/
func (s *Page) IsElementDisplayedByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByClass element display by class name*/
func (s *Page) IsElementDisplayedByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}

/*IsElementDisplayedByText element display by text in xpath*/
func (s *Page) IsElementDisplayedByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.IsDisplayed()

	return element
}
