package androidhelper

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

/*Driver global variable*/
var Driver *appium.WebDriver
var p *properties.Properties

/*AndroidProperties : read android properties content file*/
func AndroidProperties() error {
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-android.properties", properties.UTF8)

	return nil
}

/*DriverConnect for android*/
func DriverConnect() error {
	AndroidProperties()

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
