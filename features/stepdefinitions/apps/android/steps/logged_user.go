package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/support"
)

// LoggedUser : validate user has logged successfully
func LoggedUser() error {
	home := androidpages.HomePage{Page: support.AndroidApps}

	home.ClickOnboardingHomePage().
		ValidateLoggedUser()

	return nil
}
