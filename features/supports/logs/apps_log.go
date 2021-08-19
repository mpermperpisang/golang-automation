package logs

import (
	"fmt"

	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/data"
	androidformats "github.com/golang-automation/features/helper/formats/apps/android"
	iosformats "github.com/golang-automation/features/helper/formats/apps/ios"
	supports "github.com/golang-automation/features/supports/drivers"
)

func LogAndroid(log error) {
	if supports.AndroidDriver != nil {
		path := fmt.Sprintf(androidformats.AndroidPath(data.LOGS), helper.GetPWD())

		takeErrorLog(path, log)
	}
}

func LogIOS(log error) {
	if supports.IOSDriver != nil {
		path := fmt.Sprintf(iosformats.IOSPath(data.LOGS), helper.GetPWD())

		takeErrorLog(path, log)
	}
}
