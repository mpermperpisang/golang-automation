package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions/websteps"
)

/*AutomationWeb is suites for desktop or mobile web*/
func AutomationWeb(s *godog.Suite) {
	s.Step(`^client input valid data login$`, websteps.UserLogin)
	s.Step(`^client must be in logged home page$`, websteps.ValidateLoggedUser)
}
