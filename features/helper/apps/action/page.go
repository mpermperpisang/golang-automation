package action

import (
	"github.com/sclevine/agouti/appium"
)

/*Driver : Android & iOS driver*/
type Driver struct {
	Driver *appium.WebDriver
}

/*AppPage : Android & iOS Page*/
type AppPage struct {
	Page   Driver
	Device *appium.Device
}

/*Page : page actions*/
type Page struct {
	Action AppPage
}
