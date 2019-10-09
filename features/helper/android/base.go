package android

import (
	"log"

	"github.com/logrusorgru/aurora"
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

/*Driver global variable*/
var Driver *appium.WebDriver

/*Device global variable*/
var Device *appium.Device
var p *properties.Properties

/*DriverConnect for android*/
func DriverConnect() error {
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-android.properties", properties.UTF8)

	options := appium.Desired(agouti.Capabilities{
		"platformName":                p.MustGetString("platformName"),
		"platformVersion":             p.MustGetString("platformVersion"),
		"app":                         p.MustGetString("app"),
		"deviceName":                  p.MustGetString("deviceName"),
		"udid":                        p.MustGetString("udid"),
		"Debug":                       p.MustGetString("Debug"),
		"automationName":              p.MustGetString("automationName"),
		"Port":                        p.MustGetString("Port"),
		"browserName":                 p.MustGetString("browserName"),
		"noSign":                      p.MustGetString("noSign"),
		"fullReset":                   p.MustGetString("fullReset"),
		"netspeed":                    p.MustGetString("netspeed"),
		"autoGrantPermissions":        p.MustGetString("autoGrantPermissions"),
		"skipLogcatCapture":           p.MustGetString("skipLogcatCapture"),
		"disableWindowAnimation":      p.MustGetString("disableWindowAnimation"),
		"deviceReadyTimeout":          p.MustGetString("deviceReadyTimeout"),
		"espressoServerLaunchTimeout": p.MustGetString("espressoServerLaunchTimeout"),
		"Address":                     p.MustGetString("Address"),
		"appWaitActivity":             p.MustGetString("appWaitActivity"),
		"forceEspressoRebuild":        p.MustGetString("forceEspressoRebuild"),
		"uninstallOtherPackages":      p.MustGetString("uninstallOtherPackages"),
	})

	Driver = appium.New(options)

	return nil
}

/*OpenAndroidApps start and create new device android*/
func OpenAndroidApps() error {
	var err error

	if err := Driver.Start(); err != nil {
		log.Fatalln(err)
	}

	if Device, err = Driver.NewDevice(); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	return nil
}
