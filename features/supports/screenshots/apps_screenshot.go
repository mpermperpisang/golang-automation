package screenshots

import (
	"fmt"

	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/data"
	androidformat "github.com/golang-automation/features/helper/formats/apps/android"
	iosformat "github.com/golang-automation/features/helper/formats/apps/ios"
	supports "github.com/golang-automation/features/supports/drivers"
	"github.com/golang-automation/features/supports/structs"
)

func SSAndroid() {
	if supports.AndroidDriver != nil {
		path := fmt.Sprintf(androidformat.AndroidPath(data.SS), helper.GetPWD())

		takeErrorAppsImage(path)
	}
}

func SSIOS() {
	if supports.IOSDriver != nil {
		path := fmt.Sprintf(iosformat.IOSPath(data.SS), helper.GetPWD())

		takeErrorAppsImage(path)
	}
}

func takeErrorAppsImage(path string) {
	helper.DirectoryCheck(path)

	apps := structs.AppsDeviceConnect()
	filePath := helper.PathName(path, helper.FileName)

	apps.TakeScreenshot(filePath)
}
