package androidsteps

import (
	androidaction "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/stepdefinitions"
)

/*UserLogin : fill in and process login form*/
func UserLogin() error {
	action := androidaction.Page{Action: stepdefinitions.AndroidPage}
	loginProcess := androidpages.InputPhonePage{Page: action}

	loginProcess.InputPhone().
		ClickLanjut().
		InputPassword().
		ClickMasuk()

	return nil
}
