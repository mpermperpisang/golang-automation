package drivers

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

var IOSDriver *appium.WebDriver

func iOSProperties() *properties.Properties {
	return properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-ios.properties", properties.UTF8)
}

func IOSCapabilities() {
	p := iOSProperties()

	options := appium.Desired(agouti.Capabilities{
		"platformName":      p.MustGetString("platformName"),
		"deviceName":        p.MustGetString("deviceName"),
		"platformVersion":   p.MustGetString("platformVersion"),
		"app":               p.MustGetString("app"),
		"automationName":    p.MustGetString("automationName"),
		"newCommandTimeout": p.GetString("newCommandTimeout", ""),
		"fullReset":         p.GetString("fullReset", ""),
		"autoAcceptAlerts":  p.GetString("autoAcceptAlerts", ""),
		"xcodeOrgId":        p.GetString("xcodeOrgId", ""),
		"xcodeSigningId":    p.GetString("xcodeSigningId", ""),
		"udid":              p.GetString("udid", ""),
		"xcodeConfigfile":   p.GetString("xcodeConfigfile", ""),
		"useNewWDA":         p.GetString("useNewWDA", ""),
		"showXcodeLog":      p.GetString("showXcodeLog", ""),
		"browserName":       p.GetString("browserName", ""),
		"Debug":             p.GetString("Debug", ""),
	})

	IOSDriver = appium.New(options)
}
