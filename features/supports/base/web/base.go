package base

import (
	"os"

	"github.com/golang-automation-v1/features/helper"
	"github.com/golang-automation-v1/features/helper/data"
	messages "github.com/golang-automation-v1/features/helper/messages"
	"github.com/golang-automation-v1/features/supports/structs"
)

func OpenWebURL(platform, url string) error {
	var base string

	web := structs.WebDriverConnect()

	switch platform {
	case data.DWEB:
		base = data.BASE_DWEB
	case data.MWEB:
		base = data.BASE_MWEB
	default:
		helper.LogPanicln(messages.NotExistPlatform(platform) + " " + messages.PlatformList())
	}

	return web.GoToURL(os.Getenv(base) + url)
}
