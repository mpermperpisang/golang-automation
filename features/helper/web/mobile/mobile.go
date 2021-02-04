package mobilehelper

import (
	"os"

	"github.com/tebeka/selenium"
)

// MwebPage : base page url
type MwebPage struct {
	Page selenium.WebDriver
}

// BaseURL global variable
var BaseURL string

// GoToURL : access mweb url
func (s MwebPage) GoToURL(url string) error {
	BaseURL = os.Getenv("MWEB_BASE_URL") + url
	website := s.Page.Get(BaseURL)

	return website
}
