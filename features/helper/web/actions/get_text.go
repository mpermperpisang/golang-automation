package actions

import (
	"github.com/golang-automation-v1/features/helper"

	"github.com/tebeka/selenium"
)

func (s Page) GetTextByCSS(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByID(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByXpath(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByLinkText(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByPartialLink(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByName(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByText(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByContainsText(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByTag(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.Text()
}

func (s Page) GetTextByClass(locator string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.Text()
}
