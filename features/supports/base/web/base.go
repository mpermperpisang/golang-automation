package base

import (
	"os"

	"golang-automation/features/helper"
	"golang-automation/features/helper/data"
	messages "golang-automation/features/helper/messages"
	"golang-automation/features/supports/structs"
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
