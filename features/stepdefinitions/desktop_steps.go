package stepdefinitions

import (
	"os"
	"strings"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions"
)

/*LoginDWEB is function to login into desktop web*/
func LoginDWEB(user string) error {
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

	web.DriverConnect()
	web.GoToDWEBURL(os.Getenv("URL_2"))
	web.FindElementByID(objectabstractions.FieldUsername).SendKeys(username)
	web.FindElementByID(objectabstractions.FieldPassword).SendKeys(password)
	web.FindElementByText(objectabstractions.BtnLogin).Click()

	return nil
}
