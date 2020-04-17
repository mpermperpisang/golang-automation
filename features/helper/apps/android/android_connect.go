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
		"platformName":             p.MustGetString("platformName"),
		"deviceName":               p.MustGetString("deviceName"),
		"udid":                     p.MustGetString("udid"),
		"platformVersion":          p.MustGetString("platformVersion"),
		"app":                      p.MustGetString("app"),
		"automationName":           p.MustGetString("automationName"),
		"fullReset":                p.MustGetString("fullReset"),
		"espressoBuildConfig":      p.MustGetString("espressoBuildConfig"),
		"appWaitActivity":          p.MustGetString("appWaitActivity"),
		"forceEspressoRebuild":     p.MustGetString("forceEspressoRebuild"),
		"uninstallOtherPackages":   p.MustGetString("uninstallOtherPackages"),
		"netspeed":                 p.MustGetString("netspeed"),
		"noSign":                   p.MustGetString("noSign"),
		"autoGrantPermissions":     p.MustGetString("autoGrantPermissions"),
		"screenshotpath":           p.MustGetString("screenshotpath"),
		"skipDeviceInitialization": p.MustGetString("skipDeviceInitialization"),
		"skipServerInstallation":   p.MustGetString("skipServerInstallation"),
		"ignoreUnimportantViews":   p.MustGetString("ignoreUnimportantViews"),
		"Debug":                    true,
	})

	Driver = appium.New(options)

	return nil
}
