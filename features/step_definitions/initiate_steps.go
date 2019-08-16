package step_definitions

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-automation/features/demo"
	"github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/apps"
	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"
	"github.com/golang-automation/features/helper/web"
)

var usersName, meetName string

func OpenDWEB() error {
	web.DriverConnect()
	web.GoToDWEBURL(os.Getenv("URL"))
	fmt.Println(web.Driver.CurrentURL())

	return nil
}

func OpenMWEB() error {
	web.DriverConnect()
	web.GoToMWEBURL(os.Getenv("URL"))
	fmt.Println(web.Driver.CurrentURL())

	return nil
}

func OpenAndroid() error {
	android.DriverConnect()
	apps.GoToApp()

	return nil
}

func OpenIOS() error {
	ios.DriverConnect()
	apps.GoToApp()

	return nil
}

func RequestAPI(verbose string, request string) error {
	api.RetrieveAPI(verbose, request)

	return nil
}

func ResponseAPI(response int) error {
	api.ResponseAPI(response)

	return nil
}

func GivenUserName(name string) error {
	usersName = name

	return nil
}

func MeetUserName() error {
	meetName = demo.Hello(usersName)

	return nil
}

func SayHelloName(greet string) error {
	if meetName != greet {
		log.Fatalf("hello(\"Banana\") failed, expected %v, got %v", "Hello Dude!", meetName)
	}

	return nil
}
