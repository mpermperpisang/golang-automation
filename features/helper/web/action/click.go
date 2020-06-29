package webaction

import (
	"strconv"

	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// ClickByCSS : click element by CSS selector
func (s *Page) ClickByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByID : click element by class ID
func (s *Page) ClickByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByXpath : click element by Xpath selector
func (s *Page) ClickByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByLinkText : click element by link text
func (s *Page) ClickByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByPartialLink : click element by partial link text
func (s *Page) ClickByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByName : click element by class name
func (s *Page) ClickByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByTag : click element by name tag
func (s *Page) ClickByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByClass : click element by class name
func (s *Page) ClickByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	element.Click()

	return element
}

// ClickByText : click element by text in xpath
func (s *Page) ClickByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	element.Click()

	return element
}

// Numpad0 : click by numpad 0
func (s *Page) Numpad0(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad0Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad1 : click by numpad 1
func (s *Page) Numpad1(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad1Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad2 : click by numpad 2
func (s *Page) Numpad2(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad2Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad3 : click by numpad 3
func (s *Page) Numpad3(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad3Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad4 : click by numpad 4
func (s *Page) Numpad4(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad4Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad5 : click by numpad 5
func (s *Page) Numpad5(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad5Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad6 : click by numpad 6
func (s *Page) Numpad6(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad6Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad7 : click by numpad 7
func (s *Page) Numpad7(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad7Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad8 : click by numpad 8
func (s *Page) Numpad8(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad8Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}

// Numpad9 : click by numpad 9
func (s *Page) Numpad9(locator string) int {
	element, err := strconv.Atoi(selenium.Numpad9Key)
	errors.LogPanicln(err)

	s.driver().Click(element)

	return element
}
