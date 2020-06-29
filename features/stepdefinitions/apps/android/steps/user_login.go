package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

// UserLogin : fill in and process login form
func UserLogin() error {
	loginProcess := androidpages.InputPhonePage{Page: action}

	loginProcess.InputPhone().
		ClickLanjut().
		InputPassword().
		ClickMasuk()

	return nil
}
