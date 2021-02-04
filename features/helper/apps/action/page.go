package appsaction

import (
	"github.com/sclevine/agouti/appium"
)

// Page : page actions
type Page struct {
	Action AppPage
}

// AppPage : Android & iOS Page
type AppPage struct {
	Page   AppsDevice
	Driver *appium.WebDriver
}

// AppsDevice : Android & iOS device
type AppsDevice struct {
	Device *appium.Device
}

func (s Page) driver() *appium.WebDriver {
	return s.Action.Driver
}

func (s Page) device() *appium.Device {
	return s.Action.Page.Device
}
