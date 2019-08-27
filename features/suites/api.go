package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*AutomationAPI is suites for API*/
func AutomationAPI(s *godog.Suite) {
	s.Step(`^client login as "([^"]*)"$`, stepdefinitions.AuthenticationAPI)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, stepdefinitions.RequestAPIWithoutBody)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"(?: with body:)$`, stepdefinitions.RequestAPIWithBody)
	s.Step(`^response should have "([^"]*)"$`, stepdefinitions.ResponseFindKey)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, stepdefinitions.ResponseMatchingValue)
	s.Step(`^response "([^"]*)" should be (integer|string|float64)$`, stepdefinitions.ResponseDataType)
}
