package iosdriver

import (
	"fmt"
	"log"

	ios "github.com/golang-automation/features/helper/apps/ios"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device
var err error

/*StartDriver : start android driver*/
func StartDriver() error {
	if err := ios.Driver.Start(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*NewDevice : create new android device*/
func NewDevice() error {
	if Device, err = ios.Driver.NewDevice(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
