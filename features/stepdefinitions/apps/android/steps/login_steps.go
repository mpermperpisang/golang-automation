package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/support"
)

// GoToLogin : go to login page from onboarding page
func GoToLogin() error {
	goToLoginForm := androidpages.OnboardingPage{Page: support.AndroidApps}

	goToLoginForm.ClickMulaiButton().
		ClickMasuk()

	return nil
}

// UserLogin : fill in and process login form
func UserLogin() error {
	loginProcess := androidpages.InputPhonePage{Page: support.AndroidApps}

	loginProcess.InputPhone().
		ClickLanjut().
		InputPassword().
		ClickMasuk()

	return nil
}

// LoggedUser : validate user has logged successfully
func LoggedUser() error {
	home := androidpages.HomePage{Page: support.AndroidApps}

	home.ClickMengerti().
		ClickNantiSaja().
		ValidateLoggedUser()

	return nil
}
