package stepdefinitions

import (
	"os"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions"
)

/*LoginDWEB is function to login into desktop web*/
func LoginDWEB(user string) error {
	username := os.Getenv(user + "_USERNAME")
	password := os.Getenv(user + "_PASSWORD")

	web.DriverConnect()
	web.GoToDWEBURL(os.Getenv("URL_2"))
	web.FindElementByID(objectabstractions.FieldUsername).SendKeys(username)
	web.FindElementByID(objectabstractions.FieldPassword).SendKeys(password)
	web.FindElementByText(objectabstractions.BtnLogin).Click()
	web.FindElementByCSS(objectabstractions.IconProfile).Click()
	web.FindElementByText(user).IsDisplayed()

	return nil
}
