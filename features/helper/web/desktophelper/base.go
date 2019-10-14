package desktophelper

import (
	"os"

	"github.com/golang-automation/features/helper/web"
)

/*BaseURL is global variable*/
var BaseURL string

/*GoToURL is function to access dweb url*/
func GoToURL(url string) error {
	BaseURL = os.Getenv("DWEB_BASE_URL") + url
	website := web.Driver.Get(BaseURL)

	return website
}
