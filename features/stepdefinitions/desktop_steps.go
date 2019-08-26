package stepdefinitions

import (
	"log"
	"os"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions"
	"github.com/logrusorgru/aurora"
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
	if _, err := web.FindElementByText(objectabstractions.BtnSend).IsDisplayed(); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	} else {
		web.FindElementByText(objectabstractions.BtnSend).Click()
		web.FindElementByID(objectabstractions.FieldOTP).SendKeys(os.Getenv("OTP"))
		web.FindElementByText(objectabstractions.BtnConfirm).Click()
	}
	web.FindElementByCSS(objectabstractions.IconProfile).Click()
	web.FindElementByText(user).IsDisplayed()

	return nil
}
