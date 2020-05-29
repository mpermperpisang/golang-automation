package mobilesteps

import mobilepages "github.com/golang-automation/features/objectabstractions/web/mobile"

/*UserLogin : input login form*/
func UserLogin() error {
	login := mobilepages.LoginPage{Page: action}

	login.InputUsername().
		InputPassword().
		ClickBtnLogin().
		ClickSend().
		InputOTP().
		ClickNext()

	return nil
}
