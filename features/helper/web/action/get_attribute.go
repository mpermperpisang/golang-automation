package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*GetAttributeByCSS element enable by CSS selector*/
func GetAttributeByCSS(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByID element enable by class ID*/
func GetAttributeByID(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByXpath element enable by Xpath selector*/
func GetAttributeByXpath(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByLinkText element enable by link text*/
func GetAttributeByLinkText(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByPartialLink element enable by partial link text*/
func GetAttributeByPartialLink(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByName element enable by name of class*/
func GetAttributeByName(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByTag element enable by name tag*/
func GetAttributeByTag(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByClass element enable by class name*/
func GetAttributeByClass(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}

/*GetAttributeByText element enable by text in xpath*/
func GetAttributeByText(locator string, attr string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.GetAttribute(attr)
	helper.LogPanicln(err)

	return nil
}
