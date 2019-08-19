package apps

import (
	"log"

	"github.com/sclevine/agouti/appium"
)

var Device *appium.Device

func ResetApplication() error {
	if err := Device.Reset(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

func InstallApplication(app string) error {
	if err := Device.InstallApp(app); err != nil {
		log.Fatalln(err)
	}

	return nil
}

func LaunchApplication() error {
	if err := Device.LaunchApp(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

func CloseApplication() error {
	if err := Device.CloseApp(); err != nil {
		log.Fatalln(err)
	}

	return nil
}
