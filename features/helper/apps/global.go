package apps

import (
	"log"

	"github.com/sclevine/agouti/appium"
)

var Device *appium.Device

func GoToApp() error {
	if err := Device.LaunchApp(); err != nil {
		log.Fatal(err)
	}

	if err := Device.CloseApp(); err != nil {
		log.Fatal(err)
	}

	return nil
}
