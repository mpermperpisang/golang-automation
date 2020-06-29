package mobilesteps

import (
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	"github.com/golang-automation/features/support"
)

// LoggedUser : validate logged user
func LoggedUser() error {
	home := desktoppages.HomePage{Page: support.MobileWeb}

	home.ValidateLoggedUser()

	return nil
}
