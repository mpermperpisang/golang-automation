package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*SendKeysByCSS input element by CSS selector*/
func SendKeysByCSS(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByID input element by class ID*/
func SendKeysByID(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByXpath input element by Xpath selector*/
func SendKeysByXpath(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByLinkText input element by link text*/
func SendKeysByLinkText(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByPartialLink input element by partial link text*/
func SendKeysByPartialLink(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByName input element by name of class*/
func SendKeysByName(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByTag input element by name tag*/
func SendKeysByTag(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByClass input element by class name*/
func SendKeysByClass(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}

/*SendKeysByText input element by text in xpath*/
func SendKeysByText(locator string, text string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.SendKeys(text)
	helper.LogPanicln(err)

	return nil
}
