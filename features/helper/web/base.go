package web

import (
	"os"

	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func DriverConnect() error {
	caps := selenium.Capabilities{"browserName": os.Getenv("BROWSER")}
	Driver, _ = selenium.NewRemote(caps, "")

	return nil
}

func GoToURL(url string) error {
	var base string

	if os.Getenv("DWEB_BASE_URL") != "" {
		base = os.Getenv("DWEB_BASE_URL")
	} else if os.Getenv("MWEB_BASE_URL") != "" {
		base = os.Getenv("MWEB_BASE_URL")
	}

	website := Driver.Get(base + url)

	return website
}

func GoToDWEBURL(url string) error {
	website := Driver.Get(os.Getenv("DWEB_BASE_URL") + url)

	return website
}

func GoToMWEBURL(url string) error {
	website := Driver.Get(os.Getenv("MWEB_BASE_URL") + url)

	return website
}
