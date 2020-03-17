package websteps

import (
	step "github.com/golang-automation/features/stepdefinitions/web/mobile/steps"
)

/*UserLogin : input login form*/
func UserLogin() error {
	step.InputLoginForm()
	step.ClickBtnLogin()
	step.OTPConfirm()

	return nil
}

/*LoggedUser : validate logged user*/
func LoggedUser() error {
	step.ValidateLoggedUser()

	return nil
}
