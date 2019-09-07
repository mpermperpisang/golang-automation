package stepdefinitions

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"
	"github.com/golang-automation/features/helper/web"
)

var username, password string

/*LoginUI is function to set login data*/
func LoginUI(user string, platform string) error {
	readUser := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if readUser {
		username = os.Getenv(readENV + "_USERNAME")
		password = os.Getenv(readENV + "_PASSWORD")
	} else {
		username = os.Getenv(user + "_USERNAME")
		password = os.Getenv(user + "_PASSWORD")
	}

	switch platform {
	case "desktop":
		web.DriverConnect()
		web.GoToDWEBURL(os.Getenv("URL_2"))
	case "mobile":
		web.DriverConnect()
		web.GoToMWEBURL(os.Getenv("URL_2"))
	case "android":
		android.DriverConnect()
	case "ios":
		ios.DriverConnect()
	default:
		log.Fatalln("Driver not found")
	}

<<<<<<< HEAD:features/stepdefinitions/global_steps.go
=======
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

>>>>>>> cbc53593d626c3752d6bde32c0a1d95902e44096:features/stepdefinitions/desktop_steps.go
	return nil
}
