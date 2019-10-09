package stepdefinitions

import (
	"log"
	"os"

	"github.com/golang-automation/features/demo"
	"github.com/golang-automation/features/helper/android"
	"github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/ios"
	"github.com/golang-automation/features/helper/web"
)

var usersName, meetName string

/*OpenDWEB is function to initiate dweb scenario*/
func OpenDWEB() error {
	web.DriverConnect()
	web.GoToDWEBURL(os.Getenv("URL"))

	return nil
}

/*OpenMWEB is function to initiate mweb scenario*/
func OpenMWEB() error {
	web.DriverConnect()
	web.GoToMWEBURL(os.Getenv("URL"))

	return nil
}

/*OpenAndroid is function to initiate android scenario*/
func OpenAndroid() error {
	android.DriverConnect()
	android.OpenAndroidApps()

	return nil
}

/*OpenIOS is function to initiate ios scenario*/
func OpenIOS() error {
	ios.DriverConnect()
	ios.OpenIOSApps()

	return nil
}

/*BaseAPI is function to initiate base url for API*/
func BaseAPI(base string) error {
	api.BaseAPI(base)

	return nil
}

/*ResponseAPI is function to validate response code API*/
func ResponseAPI(response int) error {
	api.ResponseAPI(response)

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
	actualResult := meetName

	if expectResult := (greet); actualResult != expectResult {
		log.Fatalf("hello(\"Banana\") failed, expected %v, got %v", "Hello Dude!", meetName)
	}

	return nil
}
