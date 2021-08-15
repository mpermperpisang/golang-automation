package actions

import (
	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

func (s Page) SendKeysByCSS(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByID(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByXpath(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByLinkText(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByPartialLink(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByName(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByTag(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByClass(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByText(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.SendKeys(text)
}

func (s Page) SendKeysByContainsText(locator, text string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.SendKeys(text)
}
