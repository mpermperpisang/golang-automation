package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*IsElementDisplayedByCSS element display by CSS selector*/
func IsElementDisplayedByCSS(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByID element display by class ID*/
func IsElementDisplayedByID(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByXpath element display by Xpath selector*/
func IsElementDisplayedByXpath(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByLinkText element display by link text*/
func IsElementDisplayedByLinkText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByPartialLink element display by partial link text*/
func IsElementDisplayedByPartialLink(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByName element display by name of class*/
func IsElementDisplayedByName(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByTag element display by name tag*/
func IsElementDisplayedByTag(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByClass element display by class name*/
func IsElementDisplayedByClass(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}

/*IsElementDisplayedByText element display by text in xpath*/
func IsElementDisplayedByText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.IsDisplayed()
	helper.LogPanicln(err)

	return nil
}
