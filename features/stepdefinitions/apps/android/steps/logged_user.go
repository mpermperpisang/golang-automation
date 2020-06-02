package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*LoggedUser : validate user has logged successfully*/
func LoggedUser() error {
	home := androidpages.HomePage{Page: action}

	home.ClickOnboardingHomePage().
		ValidateLoggedUser()

	return nil
}
