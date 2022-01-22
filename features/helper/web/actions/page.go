package actions

import (
	"github.com/tebeka/selenium"
)

type Page struct {
	Action Web
}

type Web struct {
	Driver selenium.WebDriver
}

func (s Page) driver() selenium.WebDriver {
	return s.Action.Driver
}
