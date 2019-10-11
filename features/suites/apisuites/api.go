package apisuites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions/apisteps"
)

/*AutomationAPI is suites for API*/
func AutomationAPI(s *godog.Suite) {
	s.Step(`^client login as "([^"]*)"$`, apisteps.AuthenticationAPI)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, apisteps.RequestAPIWithoutBody)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"(?: with body:)$`, apisteps.RequestAPIWithBody)
	s.Step(`^response should have "([^"]*)"$`, apisteps.ResponseFindKey)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, apisteps.ResponseMatchingValue)
	s.Step(`^response "([^"]*)" should be (integer|string|float64)$`, apisteps.ResponseDataType)
}
