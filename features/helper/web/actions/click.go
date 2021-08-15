package actions

import (
	"strconv"

	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

func (s Page) ClickByCSS(locator string) error {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByID(locator string) error {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByXpath(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByLinkText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByPartialLink(locator string) error {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByName(locator string) error {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByTag(locator string) error {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByClass(locator string) error {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[text()='"+locator+"']")
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) ClickByContainsText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element.Click()
}

func (s Page) Numpad0(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad0Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad1(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad1Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad2(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad2Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad3(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad3Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad4(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad4Key)
	helper.LogPanicln(err)

	s.driver().Click(element)

	return element
}

func (s Page) Numpad5(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad5Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad6(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad6Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad7(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad7Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad8(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad8Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

func (s Page) Numpad9(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad9Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}
