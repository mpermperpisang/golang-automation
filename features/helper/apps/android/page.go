package android

import (
	"log"

	"github.com/logrusorgru/aurora"
)

/*ClickByXPath function to click xpath element*/
func ClickByXPath(locator string) error {
	element := Device.AllByXPath(locator).Click()
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return nil
}

/*ClickByText function to click text using xpath*/
func ClickByText(locator string) error {
	element := Device.FirstByXPath("//*[contains(@text, '" + locator + "')]")
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return nil
}

/*IsElementPresent function to check if element present*/
func IsElementPresent(locator string) error {
	return nil
}
