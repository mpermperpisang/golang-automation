package android

import (
	"fmt"
	"log"

	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

/*Driver global variable*/
var Driver *appium.WebDriver

/*Device global variable*/
var Device *appium.Device

/*DriverConnect for android*/
func DriverConnect() error {
	p := properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-android.properties", properties.UTF8)

	options := appium.Desired(agouti.Capabilities{
		"platformName":           p.MustGetString("platformName"),
		"deviceName":             p.MustGetString("deviceName"),
		"udid":                   p.MustGetString("udid"),
		"platformVersion":        p.MustGetString("platformVersion"),
		"browserName":            p.MustGetString("browserName"),
		"noSign":                 p.MustGetString("noSign"),
		"fullReset":              p.MustGetString("fullReset"),
		"app":                    p.MustGetString("app"),
		"netspeed":               p.MustGetString("netspeed"),
		"autoGrantPermissions":   p.MustGetString("autoGrantPermissions"),
		"automationName":         p.MustGetString("automationName"),
		"screenshotpath":         p.MustGetString("screenshotpath"),
		"appWaitActivity":        p.MustGetString("appWaitActivity"),
		"forceEspressoRebuild":   p.MustGetString("forceEspressoRebuild"),
		"uninstallOtherPackages": p.MustGetString("uninstallOtherPackages"),
	})

	Driver = appium.New(options)

	return nil
}

/*OpenApps start and create new device android*/
func OpenApps() error {
	var err error

	if err := Driver.Start(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	if Device, err = Driver.NewDevice(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
