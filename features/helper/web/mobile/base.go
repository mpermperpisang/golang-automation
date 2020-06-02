package mobilehelper

import (
	"os"

	webhelper "github.com/golang-automation/features/helper/web"
)

/*MwebPage : base page url*/
type MwebPage struct {
	Page webhelper.WebDriver
}

/*BaseURL global variable*/
var BaseURL string

/*GoToURL : access mweb url*/
func (s *MwebPage) GoToURL(url string) error {
	BaseURL = os.Getenv("MWEB_BASE_URL") + url
	website := s.Page.Driver.Get(BaseURL)

	return website
}
