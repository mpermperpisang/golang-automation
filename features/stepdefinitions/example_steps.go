package stepdefinitions

import (
	"os"

	"github.com/golang-automation/features/demo"
	"github.com/golang-automation/features/helper"
	api "github.com/golang-automation/features/helper/api"
	android "github.com/golang-automation/features/helper/apps/android"
	androidDevice "github.com/golang-automation/features/helper/apps/android/driver"
	ios "github.com/golang-automation/features/helper/apps/ios"
	iosDevice "github.com/golang-automation/features/helper/apps/ios/driver"
	"github.com/golang-automation/features/helper/message"
	web "github.com/golang-automation/features/helper/web"
	desktop "github.com/golang-automation/features/helper/web/desktop"
	mobile "github.com/golang-automation/features/helper/web/mobile"
)

var usersName, meetName string

/*OpenDWEB is function to initiate dweb scenario*/
func OpenDWEB() error {
	web.DriverConnect()
	desktop.GoToURL(os.Getenv("URL"))

	return nil
}

/*OpenMWEB is function to initiate mweb scenario*/
func OpenMWEB() error {
	web.DriverConnect()
	mobile.GoToURL(os.Getenv("URL"))

	return nil
}

/*OpenAndroid is function to initiate android scenario*/
func OpenAndroid() error {
	android.DriverConnect()
	androidDevice.StartDriver()
	androidDevice.NewDevice()

	return nil
}

/*OpenIOS is function to initiate ios scenario*/
func OpenIOS() error {
	ios.DriverConnect()
	iosDevice.StartDriver()
	iosDevice.NewDevice()

	return nil
}

/*BaseAPI is function to initiate base url for API*/
func BaseAPI(base string) error {
	api.BaseAPI(base)

	return nil
}

/*ResponseStatusAPI is function to validate response code API*/
func ResponseStatusAPI(response int) error {
	api.ResponseStatusAPI(response)

	return nil
}

/*GivenUserName is function to assign name to user*/
func GivenUserName(name string) error {
	usersName = name

	return nil
}

/*MeetUserName is function to call unit*/
func MeetUserName() error {
	meetName = demo.Hello(usersName)

	return nil
}

/*SayHelloName is function to validate unit*/
func SayHelloName(greet string) error {
	helper.AssertEqual(greet, meetName, message.UnitError())

	return nil
}
