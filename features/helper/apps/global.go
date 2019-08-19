package apps

import (
	"log"

	"github.com/sclevine/agouti/appium"
)

/*Device is a device variable for appium*/
var Device *appium.Device

/*ResetApplication is a function for reset app*/
func ResetApplication() error {
	if err := Device.Reset(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

/*InstallApplication is a function for install app*/
func InstallApplication(app string) error {
	if err := Device.InstallApp(app); err != nil {
		log.Fatalln(err)
	}

	return nil
}

/*LaunchApplication is a function for launch app*/
func LaunchApplication() error {
	if err := Device.LaunchApp(); err != nil {
		log.Fatalln(err)
	}

	return nil
}

/*CloseApplication is a function for close app*/
func CloseApplication() error {
	if err := Device.CloseApp(); err != nil {
		log.Fatalln(err)
	}

	return nil
}
