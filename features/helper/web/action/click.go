package webaction

import (
	"strconv"

	"github.com/golang-automation/features/helper"
	"github.com/tebeka/selenium"
)

// ClickByCSS : click element by CSS selector
func (s *Page) ClickByCSS(locator string) error {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByID : click element by class ID
func (s *Page) ClickByID(locator string) error {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByXpath : click element by Xpath selector
func (s *Page) ClickByXpath(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByLinkText : click element by link text
func (s *Page) ClickByLinkText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByPartialLink : click element by partial link text
func (s *Page) ClickByPartialLink(locator string) error {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByName : click element by class name
func (s *Page) ClickByName(locator string) error {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByTag : click element by name tag
func (s *Page) ClickByTag(locator string) error {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByClass : click element by class name
func (s *Page) ClickByClass(locator string) error {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element.Click()
}

// ClickByText : click element by text in xpath
func (s *Page) ClickByText(locator string) error {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(@text, '"+locator+"')]")
	helper.LogPanicln(err)

	return element.Click()
}

// Numpad0 : click by numpad 0
func (s *Page) Numpad0(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad0Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad1 : click by numpad 1
func (s *Page) Numpad1(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad1Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad2 : click by numpad 2
func (s *Page) Numpad2(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad2Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad3 : click by numpad 3
func (s *Page) Numpad3(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad3Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad4 : click by numpad 4
func (s *Page) Numpad4(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad4Key)
	helper.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad5 : click by numpad 5
func (s *Page) Numpad5(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad5Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad6 : click by numpad 6
func (s *Page) Numpad6(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad6Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad7 : click by numpad 7
func (s *Page) Numpad7(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad7Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad8 : click by numpad 8
func (s *Page) Numpad8(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad8Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}

// Numpad9 : click by numpad 9
func (s *Page) Numpad9(locator string) error {
	element, err := strconv.Atoi(selenium.Numpad9Key)
	helper.LogPanicln(err)

	return s.driver().Click(element)
}
