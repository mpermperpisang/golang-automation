package stepdefinitions

import (
	"os"
	"time"

	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions"
)

/*InputLogin is function to input username and password*/
func InputLogin() error {
	web.FindElementByID(objectabstractions.FieldUsername).SendKeys(username)
	web.FindElementByID(objectabstractions.FieldPassword).SendKeys(password)
	web.FindElementByText(objectabstractions.BtnLoginWeb).Click()
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
	web.FindElementByXpath(objectabstractions.IconProfile).IsEnabled()

	return nil
}
