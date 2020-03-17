package step

import (
	web "github.com/golang-automation/features/helper/web/action"
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	mobilepages "github.com/golang-automation/features/objectabstractions/web/mobile"
	"github.com/golang-automation/features/stepdefinitions"
)

/*InputLoginForm : input username and password*/
func InputLoginForm() error {
	web.SendKeysByID(desktoppages.FieldUsername, stepdefinitions.Username)
	web.SendKeysByID(desktoppages.FieldPassword, stepdefinitions.Password)

	return nil
}

/*ClickBtnLogin : click Login button*/
func ClickBtnLogin() error {
	web.ClickByCSS(mobilepages.BtnLogin)

	return nil
}
