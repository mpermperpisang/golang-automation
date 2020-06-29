package mobilesteps

import mobilepages "github.com/golang-automation/features/objectabstractions/web/mobile"

// UserLogin : input login form
func UserLogin() error {
	login := mobilepages.LoginPage{Page: action}

	login.InputUsername().
		InputPassword().
		ClickLogin().
		ClickSend().
		InputOTP().
		ClickNext()

	return nil
}
