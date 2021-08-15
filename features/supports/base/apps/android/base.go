package base

import (
	"github.com/golang-automation/features/supports/structs"
)

func OpenAndroidApps() error {
	startAndroidDriver()
	structs.AppsDeviceConnect()

	return nil
}

func startAndroidDriver() {
	apps := structs.AndroidDriverConnect()

	apps.StartDriver()
	apps.NewDevice()
}
