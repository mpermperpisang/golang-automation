package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*GoToLogin : go to login page from onboarding page*/
func GoToLogin() error {
	goToLoginForm := androidpages.OnboardingPage{Page: action}

	goToLoginForm.ClickBtnMulai().
		ClickBtnMasuk()

	return nil
}
