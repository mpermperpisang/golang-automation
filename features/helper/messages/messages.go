package messages

import "fmt"

func NotExistFileFormat(fileformat string) string {
	return fmt.Sprintf("Not exist file format : %s", fileformat)
}

func InvisibleElement() string {
	return "Timeout : Element is not visible"
}

func NotInteger() string {
	return "Not an integer"
}

func NotExistPlatform(platform string) string {
	return fmt.Sprintf("Not exist platform : %s", platform)
}

func PlatformList() string {
	return "Please use @api, @dweb, @mweb, @android or @ios only"
}
