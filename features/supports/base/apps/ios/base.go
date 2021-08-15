package base

import (
	"github.com/golang-automation/features/supports/structs"
)

func OpenIOSApps() error {
	startIOSDriver()
	structs.AppsDeviceConnect()

	return nil
}

func startIOSDriver() {
	apps := structs.IOSDriverConnect()

	apps.StartDriver()
	apps.NewDevice()
}
