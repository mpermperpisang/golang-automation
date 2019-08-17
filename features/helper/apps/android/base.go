package android

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
	p = properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-android.properties", properties.UTF8)

	options := appium.Desired(agouti.Capabilities{
		"platformName":    p.MustGetString("platformName"),
		"platformVersion": p.MustGetString("platformVersion"),
		"app":             p.MustGetString("app"),
		"deviceName":      p.MustGetString("deviceName"),
		"udid":            p.MustGetString("udid"),
		"Debug":           p.MustGetString("Debug"),
		"automationName":  p.MustGetString("automationName"),
		"Port":            p.MustGetString("Port"),
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
