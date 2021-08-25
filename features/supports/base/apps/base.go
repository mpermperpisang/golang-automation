package base

import (
	"github.com/mpermperpisang/golang-automation-v1/features/helper"
	"github.com/mpermperpisang/golang-automation-v1/features/helper/apps/actions"
	"github.com/mpermperpisang/golang-automation-v1/features/helper/data"
	"github.com/mpermperpisang/golang-automation-v1/features/helper/messages"
	supports "github.com/mpermperpisang/golang-automation-v1/features/supports/drivers"
	"github.com/mpermperpisang/golang-automation-v1/features/supports/structs"
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
