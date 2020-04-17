package iosdriver

import (
	"github.com/golang-automation/features/helper"
	ios "github.com/golang-automation/features/helper/apps/ios"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device

/*StartDriver : start android driver*/
func StartDriver() error {
	err := ios.Driver.Start()
	helper.LogPanicln(err)

	return nil
}

/*NewDevice : create new android device*/
func NewDevice() error {
	var err error

	Device, err = ios.Driver.NewDevice()
	helper.LogPanicln(err)

	return nil
}
