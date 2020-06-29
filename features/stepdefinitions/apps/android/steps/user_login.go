package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/support"
)

// UserLogin : fill in and process login form
func UserLogin() error {
	loginProcess := androidpages.InputPhonePage{Page: support.AndroidApps}

	loginProcess.InputPhone().
		ClickLanjut().
		InputPassword().
		ClickMasuk()

	return nil
}
