package desktopsuites

import (
	"github.com/cucumber/godog"
	desktopsteps "github.com/golang-automation/features/stepdefinitions/web/desktop/steps"
)

// AutomationDesktop : suites for desktop web
func AutomationDesktop(s *godog.ScenarioContext) {
	s.Step(`^client input valid data login$`, desktopsteps.UserLogin)
	s.Step(`^client must be in logged home page$`, desktopsteps.LoggedUser)
}
