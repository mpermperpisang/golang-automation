package desktopsteps

import (
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
)

/*LoggedUser : validate logged user*/
func LoggedUser() error {
	home := desktoppages.HomePage{Page: action}

	home.ValidateLoggedUser()

	return nil
}
