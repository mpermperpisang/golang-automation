package actions

import (
	"github.com/golang-automation-v1/features/helper"

	"github.com/tebeka/selenium"
)

func (s Page) IsElementDisplayedByCSS(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByID(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByXpath(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByLinkText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByPartialLink(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByName(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByTag(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByClass(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.IsDisplayed()
}

func (s Page) IsElementDisplayedByContainsText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.IsDisplayed()
}
