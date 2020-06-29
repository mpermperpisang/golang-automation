package mobilesteps

import (
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
