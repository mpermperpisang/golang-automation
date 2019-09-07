package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*AutomationWeb is suites for desktop or mobile web*/
func AutomationWeb(s *godog.Suite) {
	s.Step(`^client input valid data login$`, stepdefinitions.InputLogin)
	s.Step(`^client must be in logged home page$`, stepdefinitions.ValidateLoggedUser)
}
