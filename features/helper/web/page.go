package web

import (
	"github.com/tebeka/selenium"
)

/*FindElementByCSS does find element by CSS selector*/
func FindElementByCSS(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByCSSSelector, locator)
	return element
}

/*FindElementByID does find element by class ID*/
func FindElementByID(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByID, locator)
	return element
}

/*FindElementByXpath does find element by Xpath selector*/
func FindElementByXpath(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByXPATH, locator)
	return element
}

/*FindElementByLinkText does find element by link text*/
func FindElementByLinkText(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByLinkText, locator)
	return element
}

/*FindElementByPartialLink does find element by partial link text*/
func FindElementByPartialLink(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByPartialLinkText, locator)
	return element
}

/*FindElementByName does find element by name of class*/
func FindElementByName(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByName, locator)
	return element
}

/*FindElementByTag does find element by name tag*/
func FindElementByTag(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByTagName, locator)
	return element
}

/*FindElementByClass does find element by class name*/
func FindElementByClass(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByClassName, locator)
	return element
}

/*FindElementByText does find element by text in xpath*/
func FindElementByText(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	return element
}

/*FindElementsByID does find element by multiple ID*/
func FindElementsByID(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByID, locator)
	return element
}

/*FindElementsByXpath does find element by multiple Xpath*/
func FindElementsByXpath(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByXPATH, locator)
	return element
}

/*FindElementsByText does find element by multiple text using Xpath*/
func FindElementsByText(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByXPATH, "//*[contains(text(), "+locator+")]")
	return element
}

/*FindElementsByLinkText does find element by multiple link text*/
func FindElementsByLinkText(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByLinkText, locator)
	return element
}

/*FindElementsByPartialLink does find element by multiple partial link text*/
func FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByPartialLinkText, locator)
	return element
}

/*FindElementsByName does find element by multiple name of class*/
func FindElementsByName(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByName, locator)
	return element
}

/*FindElementsByTag does find element by multiple tag*/
func FindElementsByTag(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByTagName, locator)
	return element
}

/*FindElementsByClass does find element by multiple class name*/
func FindElementsByClass(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByClassName, locator)
	return element
}

/*FindElementsByCSS does find element by multiple CSS selector*/
func FindElementsByCSS(locator string) []selenium.WebElement {
	element, _ := Driver.FindElements(selenium.ByCSSSelector, locator)
	return element
}

/*MouseHoverToElement does hove to some element*/
func MouseHoverToElement(locator string) selenium.WebElement {
	element, _ := Driver.FindElement(selenium.ByCSSSelector, locator)
	element.MoveTo(0, 0)
	return element
}
