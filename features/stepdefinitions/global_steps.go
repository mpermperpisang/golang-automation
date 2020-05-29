package stepdefinitions

import (
	"fmt"
	"log"
	"os"
	"strings"

	android "github.com/golang-automation/features/helper/apps/android"
	androidhelper "github.com/golang-automation/features/helper/apps/android"
	androidDevice "github.com/golang-automation/features/helper/apps/android/driver"
	ios "github.com/golang-automation/features/helper/apps/ios"
	ioshelper "github.com/golang-automation/features/helper/apps/ios"
	iosDevice "github.com/golang-automation/features/helper/apps/ios/driver"
	"github.com/golang-automation/features/helper/data"
	"github.com/golang-automation/features/helper/message"
	web "github.com/golang-automation/features/helper/web"
	desktop "github.com/golang-automation/features/helper/web/desktop"
	mobile "github.com/golang-automation/features/helper/web/mobile"
)

/*AndroidPage : android apps page*/
var AndroidPage androidDevice.AndroidPage

/*IOSPage : ios apps page*/
var IOSPage iosDevice.IOSPage

var desktopPage desktop.DwebPage
var mobilePage mobile.MwebPage
var Dweb, Mweb web.WebDriver
var androidApps androidhelper.AndroidDriver
var iOSApps ioshelper.IOSDriver
var username string
var password string

func varLogin(user string) error {
	userENV := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if userENV {
		username = os.Getenv(readENV + "_USERNAME")
		password = os.Getenv(readENV + "_PASSWORD")
	} else {
		username = os.Getenv(user + "_USERNAME")
		password = os.Getenv(user + "_PASSWORD")
	}

	data.SetUsername(username)
	data.SetPassword(password)

	return nil
}

func dwebConnect() error {
	driver := web.DriverConnect()
	Dweb = web.WebDriver{Driver: driver}
	desktopPage = desktop.DwebPage{Page: Dweb}

	return nil
}

func mwebConnect() error {
	driver := web.DriverConnect()
	Mweb = web.WebDriver{Driver: driver}
	mobilePage = mobile.MwebPage{Page: Mweb}

	return nil
}

func androidConnectDriver() error {
	driver := android.DriverConnect()
	androidApps = android.AndroidDriver{Driver: driver}
	AndroidPage = androidDevice.AndroidPage{Page: androidApps}

	return nil
}

func iosConnectDriver() error {
	driver := ios.DriverConnect()
	iOSApps = ios.IOSDriver{Driver: driver}
	IOSPage = iosDevice.IOSPage{Page: iOSApps}

	return nil
}

func androidOpenDevice() error {
	AndroidPage.StartDriver()
	AndroidPage.NewDevice()

	return nil
}

func iosOpenDevice() error {
	IOSPage.StartDriver()
	IOSPage.NewDevice()

	return nil
}

/*LoginData : set login data*/
func LoginData(user string, platform string) error {
	varLogin(user)

	switch platform {
	case "desktop":
		dwebConnect()
		desktopPage.GoToURL(os.Getenv("URL_2"))
	case "mobile":
		mwebConnect()
		mobilePage.GoToURL(os.Getenv("URL_2"))
	case "android":
		androidConnectDriver()
		androidOpenDevice()
	case "ios":
		iosConnectDriver()
		iosOpenDevice()
	default:
		log.Panicln(fmt.Errorf(message.NotFoundDriver()))
	}

	return nil
}
