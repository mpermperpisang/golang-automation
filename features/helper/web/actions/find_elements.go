package actions

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

func (s Page) FindElementsByID(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByXpath(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByText(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByXPATH, "//*[text()="+locator+"]")
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByContainsText(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByXPATH, "//*[contains(text(), "+locator+")]")
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByLinkText(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByName(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByTag(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByClass(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element
}

func (s Page) FindElementsByCSS(locator string) []selenium.WebElement {
	element, err := s.driver().FindElements(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element
}
