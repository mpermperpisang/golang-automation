package stepdefinitions

import (
	"fmt"
	"log"
	"os"
	"strings"

	android "github.com/golang-automation/features/helper/apps/android"
	androidDriver "github.com/golang-automation/features/helper/apps/android/driver"
	ios "github.com/golang-automation/features/helper/apps/ios"
	iosDriver "github.com/golang-automation/features/helper/apps/ios/driver"
	"github.com/golang-automation/features/helper/message"
	web "github.com/golang-automation/features/helper/web"
	desktop "github.com/golang-automation/features/helper/web/desktop"
	mobile "github.com/golang-automation/features/helper/web/mobile"
)

/*Username variabel*/
var Username string

/*Password variabel*/
var Password string

func varLogin(user string) error {
	userENV := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if userENV {
		Username = os.Getenv(readENV + "_USERNAME")
		Password = os.Getenv(readENV + "_PASSWORD")
	} else {
		Username = os.Getenv(user + "_USERNAME")
		Password = os.Getenv(user + "_PASSWORD")
	}

	return nil
}

/*LoginData : set login data*/
func LoginData(user string, platform string) error {
	varLogin(user)

	switch platform {
	case "desktop":
		web.DriverConnect()
		desktop.GoToURL(os.Getenv("URL_2"))
	case "mobile":
		web.DriverConnect()
		mobile.GoToURL(os.Getenv("URL_2"))
	case "android":
		android.DriverConnect()
		androidDriver.StartDriver()
		androidDriver.NewDevice()
	case "ios":
		ios.DriverConnect()
		iosDriver.StartDriver()
		iosDriver.NewDevice()
	default:
		log.Panicln(fmt.Errorf(message.NotFoundDriver()))
	}

	return nil
}
