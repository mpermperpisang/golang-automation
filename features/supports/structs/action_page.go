package structs

import (
	appsaction "github.com/golang-automation/features/helper/apps/actions"
	webaction "github.com/golang-automation/features/helper/web/actions"
	supports "github.com/golang-automation/features/supports/drivers"
)

func WebDriverConnect() webaction.Page {
	return webaction.Page{
		Action: webaction.Web{
			Driver: supports.WebDriver,
		},
	}
}

func AndroidDriverConnect() appsaction.Page {
	return appsaction.Page{
		Action: appsaction.Apps{
			Driver: supports.AndroidDriver,
		},
	}
}

func IOSDriverConnect() appsaction.Page {
	return appsaction.Page{
		Action: appsaction.Apps{
			Driver: supports.IOSDriver,
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
