package android

import (
	"fmt"
	"log"
)

/*ClickByXPath function to click xpath element*/
func ClickByXPath(locator string) error {
	element := Device.AllByXPath(locator).Click()
	if element != nil {
		log.Panicln(fmt.Errorf("Reason: %s", element))
	}

	return nil
}

/*ClickByText function to click text using xpath*/
func ClickByText(locator string) error {
	element := Device.AllByXPath("//*[contains(@text, '" + locator + "')]").Click()
	if element != nil {
		log.Panicln(fmt.Errorf("Reason: %s", element))
	}

	return nil
}

/*InputText function to input value in textfield*/
func InputText(locator string, text string) error {
	element := Device.FindByXPath(locator).SendKeys(text)
	if element != nil {
		log.Panicln(fmt.Errorf("Reason: %s", element))
	}

	return nil
}

/*IsElementPresent function to check if element is present*/
func IsElementPresent(locator string) error {
	_, err := Device.FirstByXPath(locator).Visible()
	if err != nil {
		log.Panicln(fmt.Errorf("Reason: %s", err))
	}

	return nil
}
