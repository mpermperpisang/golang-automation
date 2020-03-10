package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*IsElementEnabledByCSS element enable by CSS selector*/
func IsElementEnabledByCSS(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByID element enable by class ID*/
func IsElementEnabledByID(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByXpath element enable by Xpath selector*/
func IsElementEnabledByXpath(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByLinkText element enable by link text*/
func IsElementEnabledByLinkText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByPartialLink element enable by partial link text*/
func IsElementEnabledByPartialLink(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByName element enable by name of class*/
func IsElementEnabledByName(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByTag element enable by name tag*/
func IsElementEnabledByTag(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByClass element enable by class name*/
func IsElementEnabledByClass(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}

/*IsElementEnabledByText element enable by text in xpath*/
func IsElementEnabledByText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.IsEnabled()
	helper.LogPanicln(err)

	return nil
}
