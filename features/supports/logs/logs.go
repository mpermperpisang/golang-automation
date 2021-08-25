package logs

import (
	"fmt"
	"io/ioutil"

	"github.com/mpermperpisang/golang-automation-v1/features/helper"
)

func takeErrorLog(path string, log error) {
	helper.DirectoryCheck(path)

	content := []byte(fmt.Sprintf("%v", log))
	filePath := helper.PathName(path, helper.FileName)

	ioutil.WriteFile(filePath, content, 0644)
}
