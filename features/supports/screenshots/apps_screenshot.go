package screenshots

import (
	"fmt"

	"golang-automation/features/helper"
	"golang-automation/features/helper/data"
	formats "golang-automation/features/helper/formats/apps"
	supports "golang-automation/features/supports/drivers"
	"golang-automation/features/supports/structs"
)

func SSAndroid() {
	if supports.AndroidDriver != nil {
		path := fmt.Sprintf(formats.AppsPath(data.SS, data.ANDROID), helper.GetPWD())

		takeErrorAppsImage(path)
	}
}

func SSIOS() {
	if supports.IOSDriver != nil {
		path := fmt.Sprintf(formats.AppsPath(data.SS, data.IOS), helper.GetPWD())

		takeErrorAppsImage(path)
	}
}

func takeErrorAppsImage(path string) {
	helper.DirectoryCheck(path)

	apps := structs.AppsDeviceConnect()
	filePath := helper.PathName(path, helper.FileName)

	apps.TakeScreenshot(filePath)
}
