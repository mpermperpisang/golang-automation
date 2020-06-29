package desktopsteps

import (
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	"github.com/golang-automation/features/support"
)

// UserLogin : input login form
func UserLogin() error {
	login := desktoppages.LoginPage{Page: support.DesktopWeb}

	login.InputUsername().
		InputPassword().
		ClickLogin().
		ClickSend().
		InputOTP().
		ClickConfirm().
		ClickOK()

	return nil
}
