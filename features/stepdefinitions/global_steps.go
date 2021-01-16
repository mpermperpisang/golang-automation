package stepdefinitions

import (
	"os"
	"strings"

	"github.com/golang-automation/features/helper/data"
	"github.com/golang-automation/features/support"
)

func varLogin(user string) error {
	var username, password string

	userENV := strings.HasPrefix(user, "ENV:")
	readENV := strings.TrimPrefix(user, "ENV:")

	if userENV {
		username = os.Getenv(readENV + "_USERNAME")
		password = os.Getenv(readENV + "_PASSWORD")
	} else {
		username = os.Getenv(user + "_USERNAME")
		password = os.Getenv(user + "_PASSWORD")
	}

	data.SetUsername(username)
	data.SetPassword(password)

	return nil
}

// LoginData : set login data
func LoginData(user string, platform string) error {
	varLogin(user)

	switch platform {
	case "desktop":
		support.DesktopPage.GoToURL(os.Getenv("URL_2"))
	case "mobile":
		support.MobilePage.GoToURL(os.Getenv("URL_2"))
	default:
		return support.AppsDevice.URLNavigate(os.Getenv("URL_2"))
	}

	return nil
}
