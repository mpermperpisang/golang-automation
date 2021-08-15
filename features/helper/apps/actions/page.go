package actions

import (
	"github.com/sclevine/agouti/appium"
)

type Page struct {
	Action Apps
}

type Apps struct {
	Driver *appium.WebDriver
	Device *appium.Device
}

func (s Page) driver() *appium.WebDriver {
	return s.Action.Driver
}

func (s Page) device() *appium.Device {
	return s.Action.Device
}
