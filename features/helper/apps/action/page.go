package appsaction

import (
	"github.com/sclevine/agouti/appium"
)

// AppsDevice : Android & iOS device
type AppsDevice struct {
	Device *appium.Device
}

// AppPage : Android & iOS Page
type AppPage struct {
	Page   AppsDevice
	Driver *appium.WebDriver
}

// Page : page actions
type Page struct {
	Action AppPage
}

func (s *Page) device() *appium.Device {
	return s.Action.Page.Device
}

func (s *Page) driver() *appium.WebDriver {
	return s.Action.Driver
}
