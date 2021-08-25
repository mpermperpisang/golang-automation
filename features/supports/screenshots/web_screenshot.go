package screenshots

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang-automation/features/helper"
	"golang-automation/features/helper/data"
	formats "golang-automation/features/helper/formats/web"
	supports "golang-automation/features/supports/drivers"
	"golang-automation/features/supports/structs"
)

func SSWeb(platform string) {
	if supports.WebDriver != nil {
		path := fmt.Sprintf(formats.WebPath(data.SS), helper.GetPWD(), strings.ToLower(platform))

		takeErrorWebImage(path)
	}
}

func takeErrorWebImage(path string) {
	helper.DirectoryCheck(path)

	web := structs.WebDriverConnect()
	filePath := helper.PathName(path, helper.FileName)

	ioutil.WriteFile(filePath, web.TakeScreenshot(), 0644)
}
