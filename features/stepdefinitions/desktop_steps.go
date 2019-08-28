package stepdefinitions

import (
	"os"
	"strings"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions"
)

/*LoginWEB is function to login into desktop and mobile web*/
func LoginWEB(user string, platform string) error {
	var username, password string

	readUser := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if readUser {
		username = os.Getenv(readENV + "_USERNAME")
		password = os.Getenv(readENV + "_PASSWORD")
	} else {
		username = os.Getenv(user + "_USERNAME")
		password = os.Getenv(user + "_PASSWORD")
	}

	if platform == "desktop" {
		web.GoToDWEBURL(os.Getenv("URL_2"))
	} else {
		web.GoToMWEBURL(os.Getenv("URL_2"))
	}

	web.DriverConnect()
	web.FindElementByID(objectabstractions.FieldUsername).SendKeys(username)
	web.FindElementByID(objectabstractions.FieldPassword).SendKeys(password)
	web.FindElementByText(objectabstractions.BtnLogin).Click()

	return nil
}
