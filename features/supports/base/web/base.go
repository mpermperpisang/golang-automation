package base

import (
	"os"

	"github.com/golang-automation/features/helper"
	messages "github.com/golang-automation/features/helper/messages/web/desktop"
	"github.com/golang-automation/features/supports/structs"
)

func OpenWebURL(platform, url string) error {
	var base string

	web := structs.WebDriverConnect()

	switch platform {
	case "dweb":
		base = "BASE_DWEB_URL"
	case "mweb":
		base = "BASE_MWEB_URL"
	default:
		helper.LogPanicln(messages.NotExistPlatform(platform))
	}

	return web.GoToURL(baseURL(base) + url)
}

func baseURL(base string) string {
	return os.Getenv(base)
}
