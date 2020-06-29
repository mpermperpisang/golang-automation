package webaction

import (
	"github.com/tebeka/selenium"
)

// Driver : Desktop & Mobile driver
type Driver struct {
	Driver selenium.WebDriver
}

// Page : page actions
type Page struct {
	Action Driver
}

func (s *Page) driver() selenium.WebDriver {
	return s.Action.Driver
}
