package webaction

import (
	"time"

	"github.com/golang-automation/features/helper/errors"
	"github.com/tebeka/selenium"
)

// FindElementByCSS : find element by CSS selector
func (s *Page) FindElementByCSS(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByID : find element by class ID
func (s *Page) FindElementByID(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByID, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByXpath : find element by Xpath selector
func (s *Page) FindElementByXpath(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByLinkText : find element by link text
func (s *Page) FindElementByLinkText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByLinkText, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByPartialLink : find element by partial link text
func (s *Page) FindElementByPartialLink(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByPartialLinkText, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByName : find element by name of class
func (s *Page) FindElementByName(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByName, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByTag : find element by name tag
func (s *Page) FindElementByTag(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByTagName, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByClass : find element by class name
func (s *Page) FindElementByClass(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByClassName, locator)
	errors.LogPanicln(err)

	return element
}

// FindElementByText : find element by text in xpath
func (s *Page) FindElementByText(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	errors.LogPanicln(err)

	return element
}

// FindElementByClickScript : find element by javascript
func (s *Page) FindElementByClickScript(locator string) []byte {
	time.Sleep(time.Second * 3)
	element, err := s.driver().ExecuteScriptRaw(`$('`+locator+`')[0].click();`, nil)
	errors.LogPanicln(err)

	return element
}

// MouseHoverToElement does hove to some element
func (s *Page) MouseHoverToElement(locator string) selenium.WebElement {
	element, err := s.driver().FindElement(selenium.ByCSSSelector, locator)
	errors.LogPanicln(err)

	element.MoveTo(0, 0)

	return element
}
