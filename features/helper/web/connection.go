package webhelper

import (
	"github.com/magiconair/properties"
	"github.com/tebeka/selenium"
)

// Driver global variables
var Driver selenium.WebDriver

// DriverConnect : connect to driver
func DriverConnect() selenium.WebDriver {
	p := properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-web.properties", properties.UTF8)
	caps := selenium.Capabilities{"browserName": p.MustGetString("browserName")}
	Driver, _ = selenium.NewRemote(caps, "")

	return Driver
}
