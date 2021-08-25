package logs

import (
	"fmt"

	"golang-automation/features/helper"
	"golang-automation/features/helper/data"
	formats "golang-automation/features/helper/formats/apps"
	supports "golang-automation/features/supports/drivers"
)

func LogAndroid(log error) {
	if supports.AndroidDriver != nil {
		path := fmt.Sprintf(formats.AppsPath(data.LOGS, data.ANDROID), helper.GetPWD())

		takeErrorLog(path, log)
	}
}

func LogIOS(log error) {
	if supports.IOSDriver != nil {
		path := fmt.Sprintf(formats.AppsPath(data.LOGS, data.IOS), helper.GetPWD())

		takeErrorLog(path, log)
	}
}
