package stepdefinitions

import (
	"log"
	"os"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/golang-automation/features/demo"
	"github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"
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

	return nil
}

/*OpenIOS is function to initiate ios scenario*/
func OpenIOS() error {
	ios.DriverConnect()

	return nil
}

/*BaseAPI is function to initiate base url for API*/
func BaseAPI(base string) error {
	api.BaseAPI(base)

	return nil
}

/*RequestAPIWithoutBody is function to initiate request API without define body in gherkin*/
func RequestAPIWithoutBody(verbose string, request string) error {
	api.RetrieveAPI(verbose, request, "")

	return nil
}

/*RequestAPIWithBody is function to initiate request API with body*/
func RequestAPIWithBody(verbose string, request string, body *gherkin.DocString) error {
	api.RetrieveAPI(verbose, request, body.Content)

	return nil
}

/*ResponseStatusCodeAPI is function to validate response API*/
func ResponseStatusCodeAPI(response int) error {
	api.ResponseStatusCodeAPI(response)

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
