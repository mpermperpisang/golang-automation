package actions

import (
	"github.com/golang-automation-v1/features/helper"

	"github.com/sclevine/agouti/appium"
)

var Device *appium.Device

func (s Page) StartDriver() {
	s.driver().Start()
}

func (s Page) NewDevice() {
	var err error

	Device, err = s.driver().NewDevice()
	helper.LogPanicln(err)
}
