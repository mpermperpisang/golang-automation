package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*FindElementsByID does find element by multiple ID*/
func (s *Page) FindElementsByID(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByXpath does find element by multiple Xpath*/
func (s *Page) FindElementsByXpath(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByText does find element by multiple text using Xpath*/
func (s *Page) FindElementsByText(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByXPATH, "//*[contains(text(), "+locator+")]")
	helper.LogPanicln(err)

	return element
}

/*FindElementsByLinkText does find element by multiple link text*/
func (s *Page) FindElementsByLinkText(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByPartialLink does find element by multiple partial link text*/
func (s *Page) FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByName does find element by multiple name of class*/
func (s *Page) FindElementsByName(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByTag does find element by multiple tag*/
func (s *Page) FindElementsByTag(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByClass does find element by multiple class name*/
func (s *Page) FindElementsByClass(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementsByCSS does find element by multiple CSS selector*/
func (s *Page) FindElementsByCSS(locator string) []selenium.WebElement {
	element, err := s.Action.Driver.FindElements(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element
}
