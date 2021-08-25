package base

import (
	"golang-automation/features/helper"
	"golang-automation/features/helper/apps/actions"
	"golang-automation/features/helper/data"
	"golang-automation/features/helper/messages"
	supports "golang-automation/features/supports/drivers"
	"golang-automation/features/supports/structs"
)

func OpenApps(platform string) {
	switch platform {
	case data.ANDROID:
		startAppsDriver(structs.AppsDriverConnect(supports.AndroidDriver))
	case data.IOS:
		startAppsDriver(structs.AppsDriverConnect(supports.IOSDriver))
	default:
		helper.LogPanicln(messages.NotExistPlatform(platform) + " " + messages.PlatformList())
	}

	structs.AppsDeviceConnect()
}

func startAppsDriver(apps actions.Page) {
	apps.StartDriver()
	apps.NewDevice()
}
