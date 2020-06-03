package desktophelper

import (
	"os"

	webhelper "github.com/golang-automation/features/helper/web"
	"gitlab.com/bot-service/helper"
)

/*DwebPage : base page url*/
type DwebPage struct {
	Page webhelper.WebDriver
}

/*BaseURL global variable*/
var BaseURL string

/*GoToURL : access dweb url*/
func (s *DwebPage) GoToURL(url string) error {
	BaseURL = os.Getenv("DWEB_BASE_URL") + url
	website := s.Page.Driver.Get(BaseURL)

	return website
}

/*ValidateURL : verify url is correct*/
func (s *DwebPage) ValidateURL(url string) error {
	actual, err := s.Page.Driver.CurrentURL()
	helper.LogPanicln(err)

	helper.AssertEqual(url, actual, "URL tidak sama")

	return nil
}
