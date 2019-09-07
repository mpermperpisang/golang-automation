package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*AutomationDWEB is suites for desktop web*/
func AutomationDWEB(s *godog.Suite) {
	s.Step(`^there is client who wants to login as "([^"]*)" via (desktop|mobile)$`, stepdefinitions.LoginWEB)
	s.Step(`^client input valid data login$`, stepdefinitions.InputLogin)
	s.Step(`^client must be in logged home page$`, stepdefinitions.ValidateLoggedUser)
}
