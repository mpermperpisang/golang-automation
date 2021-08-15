package screenshots

import (
	"fmt"

	"github.com/golang-automation/features/helper"
	androidformat "github.com/golang-automation/features/helper/formats/apps/android"
	iosformat "github.com/golang-automation/features/helper/formats/apps/ios"
	supports "github.com/golang-automation/features/supports/drivers"
	"github.com/golang-automation/features/supports/structs"
)

func SSAndroid() {
	if supports.AndroidDriver != nil {
		path := fmt.Sprintf(androidformat.SSAndroidPath(), helper.GetPWD())

		helper.DirectoryCheck(path)
		takeErrorAppsImage(path)
	}
}

func SSIOS() {
	if supports.IOSDriver != nil {
		path := fmt.Sprintf(iosformat.SSIOSPath(), helper.GetPWD())

		helper.DirectoryCheck(path)
		takeErrorAppsImage(path)
	}
}

func takeErrorAppsImage(path string) {
	apps := structs.AppsDeviceConnect()
	filePath := path + "/" + helper.FileName

	apps.TakeScreenshot(filePath)
}
