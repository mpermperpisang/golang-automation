package apisuites

import (
	"github.com/cucumber/godog"
	apihelper "github.com/golang-automation/features/helper/api"
	apisteps "github.com/golang-automation/features/stepdefinitions/api/steps"
)

// AutomationAPI : suites for API
func AutomationAPI(s *godog.ScenarioContext) {
	s.Step(`^client login as "([^"]*)"$`, apihelper.AuthenticationAPI)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, apisteps.RequestAPIWithoutBody)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"(?: with body:)?$`, apisteps.RequestAPIWithBody)
	s.Step(`^response should have "([^"]*)"$`, apisteps.ResponseFindPath)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, apisteps.ResponseMatchingValue)
	s.Step(`^response "([^"]*)" should be (integer|string|float64)$`, apisteps.ResponseDataType)
}
