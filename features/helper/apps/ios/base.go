package ios

import (
	"log"

	"github.com/magiconair/properties"
	"github.com/sclevine/agouti"
	"github.com/sclevine/agouti/appium"
)

var Driver *appium.WebDriver
var Device *appium.Device
var p *properties.Properties

func DriverConnect() error {
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-ios.properties", properties.UTF8)

	options := appium.Desired(agouti.Capabilities{
		"automationName":    p.MustGetString("automationName"),
		"platformName":      p.MustGetString("platformName"),
		"newCommandTimeout": p.MustGetString("newCommandTimeout"),
		"app":               p.MustGetString("app"),
		"platformVersion":   p.MustGetString("platformVersion"),
		"deviceName":        p.MustGetString("deviceName"),
		"fullReset":         p.MustGetString("fullReset"),
		"autoAcceptAlerts":  p.MustGetString("autoAcceptAlerts"),
		"xcodeOrgId":        p.MustGetString("xcodeOrgId"),
		"xcodeSigningId":    p.MustGetString("xcodeSigningId"),
		"udid":              p.MustGetString("udid"),
		"xcodeConfigfile":   p.MustGetString("xcodeConfigfile"),
		"useNewWDA":         p.MustGetString("useNewWDA"),
	})

	Driver = appium.New(options)

	if err := Driver.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := Driver.NewDevice(); err != nil {
		log.Fatal(err)
	}

	if err := Device.Reset(); err != nil {
		log.Fatalln(err)
	}

	if err := Device.InstallApp("{{.app}}"); err != nil {
		log.Fatalln(err)
	}

	return nil
}
