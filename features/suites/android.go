package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions/androidsteps"
)

/*AutomationApps is suites for android or ios*/
func AutomationApps(s *godog.Suite) {
	s.Step(`^client go to login page in android$`, androidsteps.GoToLogin)
	s.Step(`^client input valid data login in android$`, androidsteps.UserLogin)
	s.Step(`^client must be in logged apps page$`, androidsteps.ValidateAndroidLoggedUser)
}
