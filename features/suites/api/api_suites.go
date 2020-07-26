package apisuites

import (
	"github.com/cucumber/godog"
	apisteps "github.com/golang-automation/features/stepdefinitions/api"
)

// AutomationAPI : suites for API
func AutomationAPI(s *godog.Suite) {
	s.Step(`^client login as "([^"]*)"$`, apisteps.AuthenticationAPI)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, apisteps.RequestAPIWithoutBody)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"(?: with body:)?$`, apisteps.RequestAPIWithBody)
	s.Step(`^response should have "([^"]*)"$`, apisteps.ResponseFindPath)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, apisteps.ResponseMatchingValue)
	s.Step(`^response "([^"]*)" should be (integer|string|float64)$`, apisteps.ResponseDataType)
}
