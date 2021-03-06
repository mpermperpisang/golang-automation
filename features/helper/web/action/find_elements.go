package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// FindElementsByID : find element by multiple ID
func (s Page) FindElementsByID(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByXpath : find element by multiple Xpath
func (s Page) FindElementsByXpath(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByText : find element by multiple text using Xpath
func (s Page) FindElementsByText(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByXPATH, "//*[contains(@text, "+locator+")]")
	helper.LogPanicln(err)

	return element
}

// FindElementsByLinkText : find element by multiple link text
func (s Page) FindElementsByLinkText(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByPartialLink : find element by multiple partial link text
func (s Page) FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByName : find element by multiple name of class
func (s Page) FindElementsByName(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByTag : find element by multiple tag
func (s Page) FindElementsByTag(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByClass : find element by multiple class name
func (s Page) FindElementsByClass(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element
}

// FindElementsByCSS : find element by multiple CSS selector
func (s Page) FindElementsByCSS(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element
}
