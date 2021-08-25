package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang-automation/features/helper/data"
	"golang-automation/features/helper/formats"
	"golang-automation/features/helper/messages"
)

var FileName string

func GetPWD() string {
	pwd, err := os.Getwd()
	LogPanicln(err)

	return pwd
}

func HasENVPrefix(env string) bool {
	return strings.HasPrefix(env, "ENV:")
}

func TrimENVPrefix(env string) string {
	return strings.TrimPrefix(env, "ENV:")
}

func DirectoryCheck(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
}

func PathName(path, filename string) string {
	return filepath.Join(path, filename)
}

func SetFilename(fileformat, platform, name string) {
	var format string

	switch fileformat {
	case data.SS:
		format = formats.Screenshots()
	case data.LOGS:
		format = formats.Logs()
	default:
		LogPanicln(messages.NotExistFileFormat(fileformat))
	}

	FileName = fmt.Sprintf(format, strings.ToUpper(platform), name)
}

func RemoveContent(content string) {
	os.Remove(GetPWD() + content)
}

func RemoveAllContent(content string) {
	os.RemoveAll(GetPWD() + content)
}
