package androidhelper

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

/*Driver global variable*/
var Driver *appium.WebDriver

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
