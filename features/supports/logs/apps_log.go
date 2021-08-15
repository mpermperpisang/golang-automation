package logs

import (
	"fmt"

	"github.com/golang-automation/features/helper"
	androidformats "github.com/golang-automation/features/helper/formats/apps/android"
	iosformats "github.com/golang-automation/features/helper/formats/apps/ios"
	supports "github.com/golang-automation/features/supports/drivers"
)

func LogAndroid(log error) {
	if supports.AndroidDriver != nil {
		path := fmt.Sprintf(androidformats.LogAndroidPath(), helper.GetPWD())

		helper.DirectoryCheck(path)
		takeErrorLog(path, log)
	}
}

func LogIOS(log error) {
	if supports.IOSDriver != nil {
		path := fmt.Sprintf(iosformats.LogIOSPath(), helper.GetPWD())

		helper.DirectoryCheck(path)
		takeErrorLog(path, log)
	}
}
