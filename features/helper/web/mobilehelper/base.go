package mobilehelper

import (
	"os"

	"github.com/golang-automation/features/helper/web"
)

/*BaseURL is global variable*/
var BaseURL string

/*GoToMWEBURL is function to access mweb url*/
func GoToMWEBURL(url string) error {
	BaseURL = os.Getenv("MWEB_BASE_URL") + url
	website := web.Driver.Get(BaseURL)

	return website
}
