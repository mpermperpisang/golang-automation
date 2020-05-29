package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*IsElementSelectedByCSS element enable by CSS selector*/
func (s *Page) IsElementSelectedByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByID element enable by class ID*/
func (s *Page) IsElementSelectedByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByXpath element enable by Xpath selector*/
func (s *Page) IsElementSelectedByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByLinkText element enable by link text*/
func (s *Page) IsElementSelectedByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByPartialLink element enable by partial link text*/
func (s *Page) IsElementSelectedByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByName element enable by name of class*/
func (s *Page) IsElementSelectedByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByTag element enable by name tag*/
func (s *Page) IsElementSelectedByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByClass element enable by class name*/
func (s *Page) IsElementSelectedByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}

/*IsElementSelectedByText element enable by text in xpath*/
func (s *Page) IsElementSelectedByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.IsSelected()

	return nil
}
