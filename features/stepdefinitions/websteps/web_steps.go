package websteps

import (
	"os"
	"time"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions/webpage"
	"github.com/golang-automation/features/stepdefinitions"
)

/*UserLogin is function to fill in login form*/
func UserLogin() error {
	web.FindElementByID(webpage.FieldUsername).SendKeys(stepdefinitions.Username)
	web.FindElementByID(webpage.FieldPassword).SendKeys(stepdefinitions.Password)
	web.FindElementByText(webpage.BtnLogin).Click()
	OTPConfirm()

	return nil
}

/*OTPConfirm function to input OTP if exist*/
func OTPConfirm() error {
	web.FindElementByClickScript(webpage.BtnSend)
	time.Sleep(time.Second * 1)
	web.FindElementByID(webpage.FieldOTP).SendKeys(os.Getenv("OTP"))
	web.FindElementByClickScript(webpage.BtnConfirm)
	web.FindElementByClickScript(webpage.BtnOke)

	return nil
}

/*ValidateLoggedUser is function to check if user has login successfully*/
func ValidateLoggedUser() error {
	time.Sleep(time.Second * 10)
	web.FindElementByCSS(webpage.IconProfile).IsDisplayed()

	return nil
}
