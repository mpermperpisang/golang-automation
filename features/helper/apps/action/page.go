package appsaction

import (
	"github.com/sclevine/agouti/appium"
)

// Driver : Android & iOS driver
type Driver struct {
	Driver *appium.WebDriver
}

// AppPage : Android & iOS Page
type AppPage struct {
	Page   Driver
	Device *appium.Device
}

// Page : page actions
type Page struct {
	Action AppPage
}

func (s *Page) device() *appium.Device {
	return s.Action.Device
}

func (s *Page) driver() *appium.WebDriver {
	return s.Action.Page.Driver
}
