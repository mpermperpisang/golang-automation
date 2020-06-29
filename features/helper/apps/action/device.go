package appsaction

import (
	"github.com/golang-automation/features/helper/errors"
	"github.com/sclevine/agouti/appium"
)

// Device global variable
var Device *appium.Device

// StartDriver : start android driver
func (s *Page) StartDriver() error {
	err := s.driver().Start()
	errors.LogPanicln(err)

	return nil
}

// NewDevice : create new android device
func (s *Page) NewDevice() *appium.Device {
	var err error

	Device, err = s.driver().NewDevice()
	errors.LogPanicln(err)

	return Device
}
