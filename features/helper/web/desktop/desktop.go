package desktophelper

import (
	"os"

	"github.com/tebeka/selenium"
)

// DwebPage : base page url
type DwebPage struct {
	Page selenium.WebDriver
}

// BaseURL global variable
var BaseURL string

// GoToURL : access dweb url
func (s DwebPage) GoToURL(url string) error {
	BaseURL = os.Getenv("DWEB_BASE_URL") + url
	website := s.Page.Get(BaseURL)

	return website
}
