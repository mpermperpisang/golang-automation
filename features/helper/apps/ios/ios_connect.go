package ioshelper

import (
	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

/*Driver global variable*/
var Driver *appium.WebDriver
var p *properties.Properties

/*iOSProperties : read ios properties content file*/
func iOSProperties() error {
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-ios.properties", properties.UTF8)

	return nil
}

/*DriverConnect for ios*/
func DriverConnect() error {
	iOSProperties()

	options := appium.Desired(agouti.Capabilities{
		"platformName":    p.MustGetString("platformName"),
		"deviceName":      p.MustGetString("deviceName"),
		"platformVersion": p.MustGetString("platformVersion"),
		"app":             p.MustGetString("app"),
		"automationName":  p.MustGetString("automationName"),
		"Debug":           true,
	})

	Driver = appium.New(options)

	return nil
}
