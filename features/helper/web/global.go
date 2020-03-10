package webhelper

import (
	"os"

	"github.com/tebeka/selenium"
)

/*Driver is global variable*/
var Driver selenium.WebDriver

/*DriverConnect is function for connect to driver*/
func DriverConnect() error {
	caps := selenium.Capabilities{"browserName": os.Getenv("BROWSER")}
	Driver, _ = selenium.NewRemote(caps, "")

	return nil
}
