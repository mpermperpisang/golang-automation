package desktopsteps

import (
	"os"
	"time"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions/web/desktoppages"
	"github.com/golang-automation/features/stepdefinitions/websteps"
)

/*UserLogin is function to fill in login form*/
func UserLogin() error {
	websteps.LoginForm()
	web.FindElementByText(desktoppages.BtnLogin).Click()
	OTPConfirm()

	return nil
}

/*OTPConfirm function to input OTP if exist*/
func OTPConfirm() error {
	web.FindElementByClickScript(desktoppages.BtnSend)
	time.Sleep(time.Second * 1)
	web.FindElementByID(desktoppages.FieldOTP).SendKeys(os.Getenv("OTP"))
	web.FindElementByClickScript(desktoppages.BtnConfirm)
	web.FindElementByClickScript(desktoppages.BtnOke)

	return nil
}

/*ValidateLoggedUser is function to check if user has login successfully*/
func ValidateLoggedUser() error {
	time.Sleep(time.Second * 10)
	web.FindElementByCSS(desktoppages.IconProfile).IsDisplayed()

	return nil
}
