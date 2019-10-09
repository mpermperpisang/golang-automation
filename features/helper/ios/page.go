package ios

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/sclevine/agouti"
)

/*FindElementByClass does find element by class name*/
func FindElementByClass(locator string) *agouti.Selection {
	element := Device.FindByClass(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}

/*FindElementByButton does find element by class name*/
func FindElementByButton(locator string) *agouti.Selection {
	element := Device.FindByButton(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}

/*FindElementByID does find element by class name*/
func FindElementByID(locator string) *agouti.Selection {
	element := Device.FindByID(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}

/*FindElementByLabel does find element by class name*/
func FindElementByLabel(locator string) *agouti.Selection {
	element := Device.FindByLabel(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}

/*FindElementByLink does find element by class name*/
func FindElementByLink(locator string) *agouti.Selection {
	element := Device.FindByLink(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}

/*FindElementByName does find element by class name*/
func FindElementByName(locator string) *agouti.Selection {
	element := Device.FindByName(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}

/*FindElementByXpath does find element by class name*/
func FindElementByXpath(locator string) *agouti.Selection {
	element := Device.FindByXPath(locator)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return element
}
