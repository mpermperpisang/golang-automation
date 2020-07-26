package ioshelper

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

// Driver : global variable
var Driver *appium.WebDriver
var p *properties.Properties

// iOSProperties : read ios properties content file
func iOSProperties() error {
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-ios.properties", properties.UTF8)

	return nil
}

// DriverConnect : appium capabilities for ios
func DriverConnect() *appium.WebDriver {
	iOSProperties()

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

	Driver = appium.New(options)

	return Driver
}
