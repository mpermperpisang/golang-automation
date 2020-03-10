package mobilesteps

import (
	"os"
	"time"

	web "github.com/golang-automation/features/helper/web/action"
	mobilepages "github.com/golang-automation/features/objectabstractions/web/mobile"
	websteps "github.com/golang-automation/features/stepdefinitions/web"
)

/*UserLogin is function to fill in login form*/
func UserLogin() error {
	websteps.LoginForm()
	web.FindElementByCSS(mobilepages.BtnLogin).Click()
	OTPConfirm()

	return nil
}

/*OTPConfirm function to input OTP if exist*/
func OTPConfirm() error {
	time.Sleep(time.Second * 1)
	web.FindElementByClickScript(mobilepages.BtnSend)
	time.Sleep(time.Second * 1)
	web.FindElementByID(mobilepages.FieldOTP).SendKeys(os.Getenv("OTP"))
	web.FindElementByClickScript(mobilepages.BtnLanjut)
	time.Sleep(time.Second * 1)
	web.FindElementByClickScript(mobilepages.BtnLanjut2)

	return nil
}

/*ValidateLoggedUser is function to check if user has login successfully*/
func ValidateLoggedUser() error {
	time.Sleep(time.Second * 10)
	web.FindElementByCSS(mobilepages.IconProfile).IsDisplayed()

	return nil
}
