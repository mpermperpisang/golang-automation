package drivers

import (
	"github.com/magiconair/properties"
	"github.com/mpermperpisang/golang-automation-v1/features/helper"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

var IOSDriver *appium.WebDriver

func iOSProperties() *properties.Properties {
	return properties.MustLoadFile(helper.GetPWD()+"/capabilities-ios.properties", properties.UTF8)
}

func IOSCapabilities() {
	p := iOSProperties()
	
	options := appium.Desired(agouti.Capabilities{
		"platformName":         p.MustGetString("platformName"),
		"platformVersion":      p.MustGetString("platformVersion"),
		"deviceName":           p.MustGetString("deviceName"),
		"automationName":       p.MustGetString("automationName"),
		"newCommandTimeout":    p.MustGetString("newCommandTimeout"),
		"app":                  p.MustGetString("app"),
		"fullReset":            p.GetString("fullReset", ""),
		"netspeed":             p.GetString("netspeed", ""),
		"autoGrantPermissions": p.GetString("autoGrantPermissions", ""),
		"showXcodeLog":         p.GetString("showXcodeLog", ""),
		"Debug":                true,
	})

	IOSDriver = appium.New(options)
}
