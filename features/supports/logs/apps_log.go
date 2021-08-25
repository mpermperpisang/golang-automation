package logs

import (
	"fmt"

	"github.com/mpermperpisang/golang-automation-v1/features/helper"
	"github.com/mpermperpisang/golang-automation-v1/features/helper/data"
	formats "github.com/mpermperpisang/golang-automation-v1/features/helper/formats/apps"
	supports "github.com/mpermperpisang/golang-automation-v1/features/supports/drivers"
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
