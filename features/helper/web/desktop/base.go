package desktophelper

import (
	"os"

	webhelper "github.com/golang-automation/features/helper/web"
)

/*DwebPage : base page url*/
type DwebPage struct {
	Page webhelper.WebDriver
}

/*BaseURL : global variable*/
var BaseURL string

/*GoToURL is function to access dweb url*/
func (s *DwebPage) GoToURL(url string) error {
	BaseURL = os.Getenv("DWEB_BASE_URL") + url
	website := s.Page.Driver.Get(BaseURL)

	return website
}
