package screenshots

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/mpermperpisang/golang-automation-v1/features/helper"
	"github.com/mpermperpisang/golang-automation-v1/features/helper/data"
	formats "github.com/mpermperpisang/golang-automation-v1/features/helper/formats/web"
	supports "github.com/mpermperpisang/golang-automation-v1/features/supports/drivers"
	"github.com/mpermperpisang/golang-automation-v1/features/supports/structs"
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
