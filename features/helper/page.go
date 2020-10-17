package helper

import (
	webhelper "github.com/golang-automation/features/helper/web"
)

// WaitForLoadingPage : wait until page has load completely
func WaitForLoadingPage(timeout int) error {
	var max int

	for {
		max++

		document, err := webhelper.Driver.ExecuteScript(`return window.document.readyState`, nil)
		LogPanicln(err)

		if document == "complete" && max <= timeout {
			break
		}
	}

	return nil
}
