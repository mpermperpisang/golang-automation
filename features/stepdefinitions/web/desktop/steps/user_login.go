package desktopsteps

import (
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
)

/*UserLogin : input login form*/
func UserLogin() error {
	login := desktoppages.LoginPage{Page: action}

	login.InputUsername().
		InputPassword().
		ClickBtnLogin().
		ClickSend().
		InputOTP().
		ClickConfirm().
		ClickOK()

	return nil
}
