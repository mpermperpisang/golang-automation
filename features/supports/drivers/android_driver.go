package drivers

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

var AndroidDriver *appium.WebDriver

func androidProperties() *properties.Properties {
	return properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-android.properties", properties.UTF8)
}

func AndroidCapabilities() {
	p := androidProperties()
	
	options := appium.Desired(agouti.Capabilities{
		"platformName":         p.MustGetString("platformName"),
		"platformVersion":      p.MustGetString("platformVersion"),
		"deviceName":           p.MustGetString("deviceName"),
		"app":                  p.MustGetString("app"),
		"udid":                 p.GetString("udid", ""),
		"automationName":       p.GetString("automationName", ""),
		"fullReset":            p.GetString("fullReset", ""),
		"netspeed":             p.GetString("netspeed", ""),
		"autoGrantPermissions": p.GetString("autoGrantPermissions", ""),
		"Debug":                true,
	})

	AndroidDriver = appium.New(options)
}
