package actions

import (
	"golang-automation/features/helper"

	"github.com/tebeka/selenium"
)

func (s Page) ClearByCSS(locator string) error {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByID(locator string) error {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByXpath(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByLinkText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByPartialLink(locator string) error {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByName(locator string) error {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByTag(locator string) error {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByClass(locator string) error {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.Clear()
}

func (s Page) ClearByContainsText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.Clear()
}
