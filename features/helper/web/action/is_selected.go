package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*IsElementSelectedByCSS element enable by CSS selector*/
func IsElementSelectedByCSS(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByID element enable by class ID*/
func IsElementSelectedByID(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByXpath element enable by Xpath selector*/
func IsElementSelectedByXpath(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByLinkText element enable by link text*/
func IsElementSelectedByLinkText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByPartialLink element enable by partial link text*/
func IsElementSelectedByPartialLink(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByName element enable by name of class*/
func IsElementSelectedByName(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByTag element enable by name tag*/
func IsElementSelectedByTag(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByClass element enable by class name*/
func IsElementSelectedByClass(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}

/*IsElementSelectedByText element enable by text in xpath*/
func IsElementSelectedByText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.IsSelected()
	helper.LogPanicln(err)

	return nil
}
