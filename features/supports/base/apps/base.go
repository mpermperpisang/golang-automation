package base

import (
	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/apps/actions"
	"github.com/golang-automation/features/helper/data"
	"github.com/golang-automation/features/helper/messages"
	"github.com/golang-automation/features/supports/structs"
)

func OpenApps(platform string) {
	switch platform {
	case data.ANDROID:
		startAppsDriver(structs.AndroidDriverConnect())
	case data.IOS:
		startAppsDriver(structs.IOSDriverConnect())
	default:
		helper.LogPanicln(messages.NotExistPlatform(platform) + " " + messages.PlatformList())
	}

	structs.AppsDeviceConnect()
}

func startAppsDriver(apps actions.Page) {
	apps.StartDriver()
	apps.NewDevice()
}
