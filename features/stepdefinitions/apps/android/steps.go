package androidsteps

import step "github.com/golang-automation/features/stepdefinitions/apps/android/steps"

/*GoToLogin function to go to login page from onboarding page*/
func GoToLogin() error {
	step.ClickBtnMulai()
	step.ClickBtnMasuk()

	return nil
}

/*UserLogin function to fill in and process login form*/
func UserLogin() error {
	step.InputPhone()
	step.InputPassword()

	return nil
}

/*LoggedUser is function to validate user has logged successfully*/
func LoggedUser() error {
	step.ClickOnboardingHomePage()
	step.VerifyUserIsLogged()

	return nil
}
