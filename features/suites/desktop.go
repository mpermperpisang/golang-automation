package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*AutomationDWEB is suites for desktop web*/
func AutomationDWEB(s *godog.Suite) {
	s.Step(`^client login as "([^"]*)" via desktop$`, stepdefinitions.LoginDWEB)
}
