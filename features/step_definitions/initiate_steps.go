package step_definitions

import (
	"fmt"
	"os"

	"github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/apps"
	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"
	"github.com/golang-automation/features/helper/web"
)

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
