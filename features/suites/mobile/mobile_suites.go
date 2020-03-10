package mobilesuites

import (
	"github.com/DATA-DOG/godog"
	mobilesteps "github.com/golang-automation/features/stepdefinitions/web/mobile"
)

/*AutomationMobile is suites for mobile web*/
func AutomationMobile(s *godog.Suite) {
	s.Step(`^client input valid data login in mobile$`, mobilesteps.UserLogin)
	s.Step(`^client must be in logged mobile page$`, mobilesteps.ValidateLoggedUser)
}
