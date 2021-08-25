package base

import (
	"github.com/golang-automation-v1/features/helper"
	"github.com/golang-automation-v1/features/helper/apps/actions"
	"github.com/golang-automation-v1/features/helper/data"
	"github.com/golang-automation-v1/features/helper/messages"
	supports "github.com/golang-automation-v1/features/supports/drivers"
	"github.com/golang-automation-v1/features/supports/structs"
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
