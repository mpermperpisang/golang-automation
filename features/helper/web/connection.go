package webhelper

import (
	"github.com/magiconair/properties"
	"github.com/tebeka/selenium"
)

/*Driver is global variable*/
var Driver selenium.WebDriver

/*DriverConnect is function for connect to driver*/
func DriverConnect() error {
	p := properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-web.properties", properties.UTF8)
	caps := selenium.Capabilities{"browserName": p.MustGetString("browserName")}
	Driver, _ = selenium.NewRemote(caps, "")

	return nil
}
