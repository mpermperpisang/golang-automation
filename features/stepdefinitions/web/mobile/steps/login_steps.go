package mobilesteps

import (
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	mobilepages "github.com/golang-automation/features/objectabstractions/web/mobile"
	"github.com/golang-automation/features/support"
)

// UserLogin : input login form
func UserLogin() error {
	login := mobilepages.LoginPage{Page: support.MobileWeb}

	login.InputUsername().
		InputPassword().
		ClickLogin().
		ClickSend().
		InputOTP().
		ClickNext()

	return nil
}

// LoggedUser : validate logged user
func LoggedUser() error {
	home := desktoppages.HomePage{Page: support.MobileWeb}

	home.ValidateLoggedUser()

	return nil
}
