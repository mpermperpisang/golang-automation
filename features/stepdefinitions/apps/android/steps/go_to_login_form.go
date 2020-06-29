package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/support"
)

// GoToLogin : go to login page from onboarding page
func GoToLogin() error {
	goToLoginForm := androidpages.OnboardingPage{Page: support.AndroidApps}

	goToLoginForm.ClickMulai().
		ClickMasuk()

	return nil
}
