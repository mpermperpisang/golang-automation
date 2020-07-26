package mobilesuites

import (
	"github.com/cucumber/godog"
	mobilesteps "github.com/golang-automation/features/stepdefinitions/web/mobile/steps"
)

// AutomationMobile : suites for mobile web
func AutomationMobile(s *godog.Suite) {
	s.Step(`^client input valid data login in mobile$`, mobilesteps.UserLogin)
	s.Step(`^client must be in logged mobile page$`, mobilesteps.LoggedUser)
}
