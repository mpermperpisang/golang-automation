package base

import (
	"os"

	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/data"
	messages "github.com/golang-automation/features/helper/messages"
	"github.com/golang-automation/features/supports/structs"
)

func OpenWebURL(platform, url string) error {
	var base string

	web := structs.WebDriverConnect()

	switch platform {
	case data.DWEB:
		base = "BASE_DWEB_URL"
	case data.MWEB:
		base = "BASE_MWEB_URL"
	default:
		helper.LogPanicln(messages.NotExistPlatform(platform) + " " + messages.PlatformList())
	}

	return web.GoToURL(os.Getenv(base) + url)
}
