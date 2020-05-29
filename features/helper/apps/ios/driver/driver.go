package iosdriver

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios"
	"github.com/sclevine/agouti/appium"
)

/*IOSPage : base page app*/
type IOSPage struct {
	Page   ios.IOSDriver
	Device *appium.Device
}

/*Device global variable*/
var Device *appium.Device

/*StartDriver : start android driver*/
func (s *IOSPage) StartDriver() error {
	err := s.Page.Driver.Start()
	helper.LogPanicln(err)

	return nil
}

/*NewDevice : create new android device*/
func (s *IOSPage) NewDevice() *appium.Device {
	var err error

	Device, err = s.Page.Driver.NewDevice()
	helper.LogPanicln(err)

	return Device
}
