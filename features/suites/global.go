package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*AutomationGlobal is suites for dweb/mweb/android/ios*/
func AutomationGlobal(s *godog.Suite) {
	s.Step(`^there is client who wants to login as "([^"]*)" via (desktop|mobile|android|ios)$`, stepdefinitions.LoginUI)
}
