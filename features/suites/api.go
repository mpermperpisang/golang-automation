package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/step_definitions"
)

func AutomationAPI(s *godog.Suite) {
	s.Step(`^response should have "([^"]*)"$`, step_definitions.ResponseFindKey)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, step_definitions.ResponseMatchingValue)
}
