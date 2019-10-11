package desktopsuites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions/websteps/desktopsteps"
)

/*AutomationDesktop is suites for desktop web*/
func AutomationDesktop(s *godog.Suite) {
	s.Step(`^client input valid data login$`, desktopsteps.UserLogin)
	s.Step(`^client must be in logged home page$`, desktopsteps.ValidateLoggedUser)
}
