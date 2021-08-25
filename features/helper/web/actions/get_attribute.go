package actions

import (
	"github.com/golang-automation-v1/features/helper"

	"github.com/tebeka/selenium"
)

func (s Page) GetAttributeByCSS(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByID(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByXpath(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByLinkText(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByPartialLink(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByName(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByTag(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByClass(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByText(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}

func (s Page) GetAttributeByContainsText(locator, attr string) (string, error) {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.GetAttribute(attr)
}
