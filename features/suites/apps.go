package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*AutomationApps is suites for android or ios*/
func AutomationApps(s *godog.Suite) {
	s.Step(`^client go to login page in android$`, stepdefinitions.GoToAndroidLogin)
	s.Step(`^client input valid data login in android$`, stepdefinitions.InputAndroidLogin)
	s.Step(`^client must be in logged apps page$`, stepdefinitions.ValidateAndroidLoggedUser)
}
