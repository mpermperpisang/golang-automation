package mobilesuites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions/websteps/mobilesteps"
)

/*AutomationMobile is suites for mobile web*/
func AutomationMobile(s *godog.Suite) {
	s.Step(`^client input valid data login in mobile$`, mobilesteps.UserLogin)
	s.Step(`^client must be in logged mobile page$`, mobilesteps.ValidateLoggedUser)
}
