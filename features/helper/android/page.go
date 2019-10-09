package android

import (
	"log"
	"time"

	"github.com/logrusorgru/aurora"
)

/*ClickByXPath function to click xpath element*/
func ClickByXPath(locator string) error {
	time.Sleep(time.Second * 3)
	element := Device.AllByXPath(locator).Click()
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return nil
}

/*ClickByText function to click text using xpath*/
func ClickByText(locator string) error {
	element := Device.AllByXPath("//*[contains(@text, '" + locator + "')]").Click()
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return nil
}

/*InputText function to input value in textfield*/
func InputText(locator string, text string) error {
	element := Device.FindByXPath(locator).SendKeys(text)
	if element != nil {
		log.Fatalln(aurora.Bold(aurora.Red(element)))
	}

	return nil
}

/*IsElementPresent function to check if element is present*/
func IsElementPresent(locator string) error {
	_, err := Device.FirstByXPath(locator).Visible()
	if err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	return nil
}
