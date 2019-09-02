package web

import (
	"log"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/tebeka/selenium"
)

/*FindElementByCSS does find element by CSS selector*/
func FindElementByCSS(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByCSSSelector, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	return element
}

/*FindElementByID does find element by class ID*/
func FindElementByID(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByID, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByXpath does find element by Xpath selector*/
func FindElementByXpath(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByXPATH, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByLinkText does find element by link text*/
func FindElementByLinkText(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByLinkText, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByPartialLink does find element by partial link text*/
func FindElementByPartialLink(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByPartialLinkText, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByName does find element by name of class*/
func FindElementByName(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByName, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByTag does find element by name tag*/
func FindElementByTag(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByTagName, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByClass does find element by class name*/
func FindElementByClass(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByClassName, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByText does find element by text in xpath*/
func FindElementByText(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByXPATH, "//*[contains(text(), '"+locator+"')]")
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementByClickScript does find element by javascript*/
func FindElementByClickScript(locator string) error {
	time.Sleep(time.Second * 1)
	_, err := Driver.ExecuteScriptRaw(`$('`+locator+`').click();`, nil)
	if err != nil {
		panic(err)
		// log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	return nil
}

/*FindElementsByID does find element by multiple ID*/
func FindElementsByID(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByID, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByXpath does find element by multiple Xpath*/
func FindElementsByXpath(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByXPATH, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByText does find element by multiple text using Xpath*/
func FindElementsByText(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByXPATH, "//*[contains(text(), "+locator+")]")
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByLinkText does find element by multiple link text*/
func FindElementsByLinkText(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByLinkText, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByPartialLink does find element by multiple partial link text*/
func FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByPartialLinkText, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByName does find element by multiple name of class*/
func FindElementsByName(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByName, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByTag does find element by multiple tag*/
func FindElementsByTag(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByTagName, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByClass does find element by multiple class name*/
func FindElementsByClass(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByClassName, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*FindElementsByCSS does find element by multiple CSS selector*/
func FindElementsByCSS(locator string) []selenium.WebElement {
	element, err := Driver.FindElements(selenium.ByCSSSelector, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	return element
}

/*MouseHoverToElement does hove to some element*/
func MouseHoverToElement(locator string) selenium.WebElement {
	element, err := Driver.FindElement(selenium.ByCSSSelector, locator)
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}
	element.MoveTo(0, 0)
	return element
}
