package actions

import (
	"golang-automation/features/helper"

	"github.com/tebeka/selenium"
)

func (s Page) IsElementEnabledByCSS(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByID(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByXpath(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByLinkText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByPartialLink(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByName(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByTag(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByClass(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.IsEnabled()
}

func (s Page) IsElementEnabledByContainsText(locator string) (bool, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.IsEnabled()
}
