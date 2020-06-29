package desktopsuites

import (
	"github.com/DATA-DOG/godog"
	desktopsteps "github.com/golang-automation/features/stepdefinitions/web/desktop/steps"
)

// AutomationDesktop : suites for desktop web
func AutomationDesktop(s *godog.Suite) {
	s.Step(`^client input valid data login$`, desktopsteps.UserLogin)
	s.Step(`^client must be in logged home page$`, desktopsteps.LoggedUser)
}
