package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*GetTextByCSS get text element by CSS selector*/
func (s *Page) GetTextByCSS(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByID get text element by class ID*/
func (s *Page) GetTextByID(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByXpath get text element by Xpath selector*/
func (s *Page) GetTextByXpath(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByLinkText get text element by link text*/
func (s *Page) GetTextByLinkText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByPartialLink get text element by partial link text*/
func (s *Page) GetTextByPartialLink(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByName get text element by name of class*/
func (s *Page) GetTextByName(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByText get text element by text in xpath*/
func (s *Page) GetTextByText(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByTag get text element by name tag*/
func (s *Page) GetTextByTag(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}

/*GetTextByClass get text element by class name*/
func (s *Page) GetTextByClass(locator string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.Text()

	return element
}
