package androidsteps

import (
	androidaction "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/stepdefinitions"
)

/*LoggedUser is function to validate user has logged successfully*/
func LoggedUser() error {
	action := androidaction.Page{Action: stepdefinitions.AndroidPage}
	home := androidpages.HomePage{Page: action}

	home.ClickOnboardingHomePage().
		VerifyUserIsLogged()

	return nil
}
