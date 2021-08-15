package screenshots

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/golang-automation/features/helper"
	formats "github.com/golang-automation/features/helper/formats/web"
	supports "github.com/golang-automation/features/supports/drivers"
	"github.com/golang-automation/features/supports/structs"
)

func SSWeb(platform string) {
	if supports.WebDriver != nil {
		path := fmt.Sprintf(formats.SSWebPath(), helper.GetPWD(), strings.ToLower(platform))

		helper.DirectoryCheck(path)
		takeErrorWebImage(path)
	}
}

func takeErrorWebImage(path string) {
	web := structs.WebDriverConnect()
	filePath := path + "/" + helper.FileName
	err := ioutil.WriteFile(filePath, web.TakeScreenshot(), 0644)
	helper.LogPanicln(err)
}
