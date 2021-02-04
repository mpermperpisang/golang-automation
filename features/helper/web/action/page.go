package webaction

import (
	"github.com/tebeka/selenium"
)

// Page : page actions
type Page struct {
	Action selenium.WebDriver
}

func (s Page) driver() selenium.WebDriver {
	return s.Action
}
