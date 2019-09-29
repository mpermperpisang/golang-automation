package web

import (
	"os"

	"github.com/tebeka/selenium"
)

/*Driver is global variable*/
var Driver selenium.WebDriver

/*BaseURL is global variable*/
var BaseURL string

/*DriverConnect is function for connect to driver*/
func DriverConnect() error {
	caps := selenium.Capabilities{"browserName": os.Getenv("BROWSER")}
	Driver, _ = selenium.NewRemote(caps, "")

	return nil
}

/*GoToDWEBURL is function to access dweb url*/
func GoToDWEBURL(url string) error {
	BaseURL = os.Getenv("DWEB_BASE_URL") + url
	website := Driver.Get(BaseURL)

	return website
}

/*GoToMWEBURL is function to access mweb url*/
func GoToMWEBURL(url string) error {
	BaseURL = os.Getenv("MWEB_BASE_URL") + url
	website := Driver.Get(BaseURL)

	return website
}
