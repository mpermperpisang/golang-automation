package webaction

import (
	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*GetTextByCSS get text element by CSS selector*/
func GetTextByCSS(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByID get text element by class ID*/
func GetTextByID(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByXpath get text element by Xpath selector*/
func GetTextByXpath(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByLinkText get text element by link text*/
func GetTextByLinkText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByPartialLink get text element by partial link text*/
func GetTextByPartialLink(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByName get text element by name of class*/
func GetTextByName(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByText get text element by text in xpath*/
func GetTextByText(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByTag get text element by name tag*/
func GetTextByTag(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}

/*GetTextByClass get text element by class name*/
func GetTextByClass(locator string) error {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	element.Text()
	helper.LogPanicln(err)

	return nil
}
