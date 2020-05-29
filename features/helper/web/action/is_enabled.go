package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*IsElementEnabledByCSS element enable by CSS selector*/
func (s *Page) IsElementEnabledByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByID element enable by class ID*/
func (s *Page) IsElementEnabledByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByXpath element enable by Xpath selector*/
func (s *Page) IsElementEnabledByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByLinkText element enable by link text*/
func (s *Page) IsElementEnabledByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByPartialLink element enable by partial link text*/
func (s *Page) IsElementEnabledByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByName element enable by name of class*/
func (s *Page) IsElementEnabledByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByTag element enable by name tag*/
func (s *Page) IsElementEnabledByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByClass element enable by class name*/
func (s *Page) IsElementEnabledByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}

/*IsElementEnabledByText element enable by text in xpath*/
func (s *Page) IsElementEnabledByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.IsEnabled()

	return element
}
