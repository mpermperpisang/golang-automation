package structs

import (
	appsaction "golang-automation/features/helper/apps/actions"
	webaction "golang-automation/features/helper/web/actions"
	supports "golang-automation/features/supports/drivers"

	"github.com/sclevine/agouti/appium"
)

func WebDriverConnect() webaction.Page {
	return webaction.Page{
		Action: webaction.Web{
			Driver: supports.WebDriver,
		},
	}
}

func AppsDriverConnect(driver *appium.WebDriver) appsaction.Page {
	return appsaction.Page{
		Action: appsaction.Apps{
			Driver: driver,
		},
	}
}

func AppsDeviceConnect() appsaction.Page {
	return appsaction.Page{
		Action: appsaction.Apps{
			Device: appsaction.Device,
		},
	}
}
