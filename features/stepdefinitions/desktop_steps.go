package stepdefinitions

import (
	"os"
	"strings"
	"time"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions"
)

var username, password string

/*LoginWEB is function to set login data*/
func LoginWEB(user string, platform string) error {
	web.DriverConnect()

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

	return nil
}

/*InputLogin is function to input username and password*/
func InputLogin() error {
	web.FindElementByID(objectabstractions.FieldUsername).SendKeys(username)
	web.FindElementByID(objectabstractions.FieldPassword).SendKeys(password)
	web.FindElementByText(objectabstractions.BtnLogin).Click()
	web.FindElementByClickScript(objectabstractions.BtnSend)
	time.Sleep(time.Second * 1)
	web.FindElementByID(objectabstractions.FieldOTP).SendKeys(os.Getenv("OTP"))
	web.FindElementByClickScript(objectabstractions.BtnConfirm)
	web.FindElementByClickScript(objectabstractions.BtnOke)

	return nil
}

/*ValidateLoggedUser is function to check if user has login successfully*/
func ValidateLoggedUser() error {
	time.Sleep(time.Second * 10)
	web.FindElementByXpath(objectabstractions.IconProfile).IsDisplayed()

	return nil
}
