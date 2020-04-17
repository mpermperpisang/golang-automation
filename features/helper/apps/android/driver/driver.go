package androiddriver

import (
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device

/*StartDriver : start android driver*/
func StartDriver() error {
	err := android.Driver.Start()
	helper.LogPanicln(err)

	return nil
}

/*NewDevice : create new android device*/
func NewDevice() error {
	var err error

	Device, err = android.Driver.NewDevice()
	helper.LogPanicln(err)

	return nil
}
