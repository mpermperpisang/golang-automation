package logs

import (
	"fmt"
	"strings"

	"github.com/golang-automation-v1/features/helper"
	"github.com/golang-automation-v1/features/helper/data"
	formats "github.com/golang-automation-v1/features/helper/formats/web"
	supports "github.com/golang-automation-v1/features/supports/drivers"
)

func LogWeb(platform string, log error) {
	if supports.WebDriver != nil {
		path := fmt.Sprintf(formats.WebPath(data.LOGS), helper.GetPWD(), strings.ToLower(platform))

		takeErrorLog(path, log)
	}
}
