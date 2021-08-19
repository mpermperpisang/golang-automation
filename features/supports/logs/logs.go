package logs

import (
	"fmt"
	"io/ioutil"

	"github.com/golang-automation/features/helper"
)

func takeErrorLog(path string, log error) {
	helper.DirectoryCheck(path)

	content := []byte(fmt.Sprintf("%v", log))
	filePath := helper.PathName(path, helper.FileName)
	err := ioutil.WriteFile(filePath, content, 0644)
	helper.LogPanicln(err)
}
