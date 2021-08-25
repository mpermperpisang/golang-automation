package screenshots

import (
	"fmt"

	"github.com/mpermperpisang/golang-automation-v1/features/helper"
	"github.com/mpermperpisang/golang-automation-v1/features/helper/data"
	formats "github.com/mpermperpisang/golang-automation-v1/features/helper/formats/apps"
	supports "github.com/mpermperpisang/golang-automation-v1/features/supports/drivers"
	"github.com/mpermperpisang/golang-automation-v1/features/supports/structs"
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
