package stepdefinitions

import (
	"log"
	"os"
	"strings"

	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"
	"github.com/golang-automation/features/helper/web"
)

var username, password string

/*LoginUI is function to set login data*/
func LoginUI(user string, platform string) error {
	userENV := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if userENV {
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
		android.OpenAndroidApps()
	case "ios":
		ios.DriverConnect()
		ios.OpenIOSApps()
	default:
		log.Fatalln("Driver not found")
	}

	return nil
}
