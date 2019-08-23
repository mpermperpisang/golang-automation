package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

func AutomationAPI(s *godog.Suite) {
	s.Step(`^client login as "([^"]*)"$`, stepdefinitions.Authentication)
	s.Step(`^response should have "([^"]*)"$`, stepdefinitions.ResponseFindKey)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, stepdefinitions.ResponseMatchingValue)
}
