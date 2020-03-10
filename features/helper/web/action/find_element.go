package webaction

import (
	"time"

	"github.com/golang-automation/features/helper"
	web "github.com/golang-automation/features/helper/web"
	"github.com/tebeka/selenium"
)

/*FindElementByCSS does find element by CSS selector*/
func FindElementByCSS(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByID does find element by class ID*/
func FindElementByID(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByID, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByXpath does find element by Xpath selector*/
func FindElementByXpath(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByXPATH, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByLinkText does find element by link text*/
func FindElementByLinkText(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByLinkText, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByPartialLink does find element by partial link text*/
func FindElementByPartialLink(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByPartialLinkText, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByName does find element by name of class*/
func FindElementByName(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByName, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByTag does find element by name tag*/
func FindElementByTag(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByTagName, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByClass does find element by class name*/
func FindElementByClass(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByClassName, locator)
	helper.LogPanicln(err)

	return element
}

/*FindElementByText does find element by text in xpath*/
func FindElementByText(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	helper.LogPanicln(err)

	return element
}

/*FindElementByClickScript does find element by javascript*/
func FindElementByClickScript(locator string) error {
	time.Sleep(time.Second * 3)
	_, err := web.Driver.ExecuteScriptRaw(`$('`+locator+`')[0].click();`, nil)
	helper.LogPanicln(err)

	return nil
}

/*MouseHoverToElement does hove to some element*/
func MouseHoverToElement(locator string) selenium.WebElement {
	element, err := web.Driver.FindElement(selenium.ByCSSSelector, locator)
	helper.LogPanicln(err)
	element.MoveTo(0, 0)

	return element
}
