package step

import (
	web "github.com/golang-automation/features/helper/web/action"
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	"github.com/golang-automation/features/stepdefinitions"
)

/*InputLoginForm : fill in username and password*/
func InputLoginForm() error {
	web.SendKeysByID(desktoppages.FieldUsername, stepdefinitions.Username)
	web.SendKeysByID(desktoppages.FieldPassword, stepdefinitions.Password)

	return nil
}

/*ClickBtnLogin : click Login button*/
func ClickBtnLogin() error {
	web.ClickByText(desktoppages.BtnLogin)

	return nil
}
