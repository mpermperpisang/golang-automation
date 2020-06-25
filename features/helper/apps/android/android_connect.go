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
func DriverConnect() *appium.WebDriver {
	AndroidProperties()

	options := appium.Desired(agouti.Capabilities{
		"platformName":             p.MustGetString("platformName"),
		"deviceName":               p.MustGetString("deviceName"),
		"udid":                     p.MustGetString("udid"),
		"platformVersion":          p.MustGetString("platformVersion"),
		"app":                      p.MustGetString("app"),
		"automationName":           p.MustGetString("automationName"),
		"fullReset":                p.GetString("fullReset", ""),
		"espressoBuildConfig":      p.GetString("espressoBuildConfig", ""),
		"appWaitActivity":          p.GetString("appWaitActivity", ""),
		"forceEspressoRebuild":     p.GetString("forceEspressoRebuild", ""),
		"uninstallOtherPackages":   p.GetString("uninstallOtherPackages", ""),
		"netspeed":                 p.GetString("netspeed", ""),
		"noSign":                   p.GetString("noSign", ""),
		"autoGrantPermissions":     p.GetString("autoGrantPermissions", ""),
		"screenshotpath":           p.GetString("screenshotpath", ""),
		"skipDeviceInitialization": p.GetString("skipDeviceInitialization", ""),
		"skipServerInstallation":   p.GetString("skipServerInstallation", ""),
		"ignoreUnimportantViews":   p.GetString("ignoreUnimportantViews", ""),
		"Debug":                    true,
	})

	Driver = appium.New(options)

	return Driver
}
