package actions

import (
	"golang-automation/features/helper"

	"github.com/tebeka/selenium"
)

func (s Page) IsElementSelectedByCSS(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByID(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByXpath(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByLinkText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByPartialLink(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByName(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByTag(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByClass(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.IsSelected()
}

func (s Page) IsElementSelectedByContainsText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.IsSelected()
}
