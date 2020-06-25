package action

import (
	"github.com/golang-automation/features/helper"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device

/*StartDriver : start android driver*/
func (s *AppPage) StartDriver() error {
	err := s.Page.Driver.Start()
	helper.LogPanicln(err)

	return nil
}

/*NewDevice : create new android device*/
func (s *AppPage) NewDevice() *appium.Device {
	var err error

	Device, err = s.Page.Driver.NewDevice()
	helper.LogPanicln(err)

	return Device
}
