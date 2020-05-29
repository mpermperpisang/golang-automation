package androidsteps

import (
	androidaction "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/stepdefinitions"
)

/*GoToLogin : go to login page from onboarding page*/
func GoToLogin() error {
	action := androidaction.Page{Action: stepdefinitions.AndroidPage}
	goToLoginForm := androidpages.OnboardingPage{Page: action}

	goToLoginForm.ClickBtnMulai().
		ClickBtnMasuk()

	return nil
}
