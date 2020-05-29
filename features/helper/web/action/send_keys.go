package webaction

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

/*SendKeysByCSS input element by CSS selector*/
func (s *Page) SendKeysByCSS(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByID input element by class ID*/
func (s *Page) SendKeysByID(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByXpath input element by Xpath selector*/
func (s *Page) SendKeysByXpath(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByLinkText input element by link text*/
func (s *Page) SendKeysByLinkText(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByPartialLink input element by partial link text*/
func (s *Page) SendKeysByPartialLink(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByName input element by name of class*/
func (s *Page) SendKeysByName(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByTag input element by name tag*/
func (s *Page) SendKeysByTag(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByClass input element by class name*/
func (s *Page) SendKeysByClass(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}

/*SendKeysByText input element by text in xpath*/
func (s *Page) SendKeysByText(locator string, text string) selenium.WebElement {
	element, err := s.Action.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	element.SendKeys(text)

	return nil
}
