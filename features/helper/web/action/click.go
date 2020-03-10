package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*ClickByCSS click element by CSS selector*/
func ClickByCSS(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByID click element by class ID*/
func ClickByID(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByXpath click element by Xpath selector*/
func ClickByXpath(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByLinkText click element by link text*/
func ClickByLinkText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByPartialLink click element by partial link text*/
func ClickByPartialLink(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByName click element by class name*/
func ClickByName(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByTag click element by name tag*/
func ClickByTag(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByClass click element by class name*/
func ClickByClass(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.Click()
	helper.LogPanicln(err)

	return nil
}

/*ClickByText click element by text in xpath*/
func ClickByText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.Click()
	helper.LogPanicln(err)

	return nil
}
