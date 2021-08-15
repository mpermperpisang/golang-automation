package logs

import (
	"fmt"
	"strings"

	"github.com/golang-automation/features/helper"
	formats "github.com/golang-automation/features/helper/formats/web"
	supports "github.com/golang-automation/features/supports/drivers"
)

func LogWeb(platform string, log error) {
	if supports.WebDriver != nil {
		path := fmt.Sprintf(formats.LogWebPath(), helper.GetPWD(), strings.ToLower(platform))

		helper.DirectoryCheck(path)
		takeErrorLog(path, log)
	}
}
