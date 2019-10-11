package websteps

import (
	"github.com/golang-automation/features/helper/web"
	"github.com/golang-automation/features/objectabstractions/web/desktoppages"
	"github.com/golang-automation/features/stepdefinitions"
)

/*LoginForm function to fill in username and password*/
func LoginForm() error {
	web.FindElementByID(desktoppages.FieldUsername).SendKeys(stepdefinitions.Username)
	web.FindElementByID(desktoppages.FieldPassword).SendKeys(stepdefinitions.Password)

	return nil
}
