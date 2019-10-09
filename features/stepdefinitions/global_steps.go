package stepdefinitions

import (
	"log"
	"os"
	"strings"

	"github.com/golang-automation/features/helper/android"
	"github.com/golang-automation/features/helper/ios"
	"github.com/golang-automation/features/helper/web"
)

/*Username variabel*/
var Username string

/*Password variabel*/
var Password string

/*LoginData is function to set login data*/
func LoginData(user string, platform string) error {
	userENV := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if userENV {
		Username = os.Getenv(readENV + "_USERNAME")
		Password = os.Getenv(readENV + "_PASSWORD")
	} else {
		Username = os.Getenv(user + "_USERNAME")
		Password = os.Getenv(user + "_PASSWORD")
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
		android.OpenAndroidApps()
	case "ios":
		ios.DriverConnect()
		ios.OpenIOSApps()
	default:
		log.Fatalln("Driver not found")
	}

	return nil
}
