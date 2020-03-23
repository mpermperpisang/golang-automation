package androiddriver

import (
	"fmt"
	"log"

	android "github.com/golang-automation/features/helper/apps/android"
	"github.com/sclevine/agouti/appium"
)

/*Device global variable*/
var Device *appium.Device
var err error

/*StartDriver : start android driver*/
func StartDriver() error {
	if err := android.Driver.Start(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*NewDevice : create new android device*/
func NewDevice() error {
	if Device, err = android.Driver.NewDevice(); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
