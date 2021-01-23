package page

import (
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/action"
	webhelper "github.com/golang-automation/features/helper/web"
)

// WaitForLoadingPage : wait until page has load completely
func WaitForLoadingPage(timeout int) error {
	var max int

	for {
		max++

		document, err := webhelper.Driver.ExecuteScript(`return window.document.readyState`, nil)
		helper.LogPanicln(err)

		if document == "complete" && max <= timeout {
			break
		}
	}

	return nil
}

// ValidateElementWithTimeout : TO DO
func ValidateElementWithTimeout(element bool, timeout int) error {
	var timeCount int = 0

	for loop := true; loop; loop = !element {
		android.Device.SetImplicitWait(timeCount)

		if timeCount <= timeout {
			timeCount++
		} else {
			break
		}
	}

	return nil
}
