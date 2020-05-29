package androiddriver

import (
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android"
	"github.com/sclevine/agouti/appium"
)

/*AndroidPage : base page apk*/
type AndroidPage struct {
	Page   android.AndroidDriver
	Device *appium.Device
}

/*Device global variable*/
var Device *appium.Device

/*StartDriver : start android driver*/
func (s *AndroidPage) StartDriver() error {
	err := s.Page.Driver.Start()
	helper.LogPanicln(err)

	return nil
}

/*NewDevice : create new android device*/
func (s *AndroidPage) NewDevice() *appium.Device {
	var err error

	Device, err = s.Page.Driver.NewDevice()
	helper.LogPanicln(err)

	return Device
}
