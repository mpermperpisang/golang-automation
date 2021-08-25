package logs

import (
	"fmt"
	"strings"

	"golang-automation/features/helper"
	"golang-automation/features/helper/data"
	formats "golang-automation/features/helper/formats/web"
	supports "golang-automation/features/supports/drivers"
)

func LogWeb(platform string, log error) {
	if supports.WebDriver != nil {
		path := fmt.Sprintf(formats.WebPath(data.LOGS), helper.GetPWD(), strings.ToLower(platform))

		takeErrorLog(path, log)
	}
}
