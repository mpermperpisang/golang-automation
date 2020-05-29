package androidsuites

import (
	"github.com/DATA-DOG/godog"
	androidsteps "github.com/golang-automation/features/stepdefinitions/apps/android/steps"
)

/*AutomationAndroid is suites for android*/
func AutomationAndroid(s *godog.Suite) {
	s.Step(`^client go to login page in android$`, androidsteps.GoToLogin)
	s.Step(`^client input valid data login in android$`, androidsteps.UserLogin)
	s.Step(`^client must be in logged android page$`, androidsteps.LoggedUser)
}
