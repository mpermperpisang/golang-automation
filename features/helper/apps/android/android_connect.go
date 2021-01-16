package androidhelper

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

// Driver : global variable
var Driver *appium.WebDriver
var p *properties.Properties

// AndroidProperties : read android properties content file
func AndroidProperties() error {
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-android.properties", properties.UTF8)

	return nil
}

// DriverConnect : appium capabilities for android
func DriverConnect() *appium.WebDriver {
	AndroidProperties()

	options := appium.Desired(agouti.Capabilities{
		"platformName":             p.MustGetString("platformName"),
		"deviceName":               p.MustGetString("deviceName"),
		"udid":                     p.MustGetString("udid"),
		"platformVersion":          p.MustGetString("platformVersion"),
		"browserName":              p.GetString("browserName", ""),
		"noSign":                   p.GetString("noSign", ""),
		"fullReset":                p.GetString("fullReset", ""),
		"app":                      p.MustGetString("app"),
		"netspeed":                 p.GetString("netspeed", ""),
		"autoGrantPermissions":     p.GetString("autoGrantPermissions", ""),
		"automationName":           p.MustGetString("automationName"),
		"screenshotpath":           p.GetString("screenshotpath", ""),
		"appWaitActivity":          p.GetString("appWaitActivity", ""),
		"forceEspressoRebuild":     p.GetString("forceEspressoRebuild", ""),
		"uninstallOtherPackages":   p.GetString("uninstallOtherPackages", ""),
		"espressoBuildConfig":      p.GetString("espressoBuildConfig", ""),
		"appWaitDuration":          p.GetString("appWaitDuration", ""),
		"skipDeviceInitialization": p.GetString("skipDeviceInitialization", ""),
		"skipServerInstallation":   p.GetString("skipServerInstallation", ""),
		"ignoreUnimportantViews":   p.GetString("ignoreUnimportantViews", ""),
		"Debug":                    true,
	})

	Driver = appium.New(options)

	return Driver
}
