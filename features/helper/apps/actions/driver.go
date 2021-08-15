package actions

import (
	"github.com/golang-automation/features/helper"
	"github.com/sclevine/agouti/appium"
)

var Device *appium.Device

func (s Page) StartDriver() error {
	return s.driver().Start()
}

func (s Page) NewDevice() *appium.Device {
	var err error

	Device, err = s.driver().NewDevice()
	helper.LogPanicln(err)

	return Device
}
