package suites

import (
	"github.com/cucumber/godog"

	stepdefinitions "github.com/golang-automation-v1/features/stepdefinitions/api"
)

func APIScenarioContext(s *godog.ScenarioContext) {
	s.Step(`client sends a ([^\"]*) to "([^\"]*)"`, stepdefinitions.HitEndpointWithoutBody)
	s.Step(`client sends a ([^\"]*) to "([^\"]*)" with body:`, stepdefinitions.HitEndpointWithBody)
	s.Step(`validate status is (\d+)`, stepdefinitions.ValidateStatus)
	s.Step(`validate "([^\"]*)" is "([^\"]*)"`, stepdefinitions.ValidateValue)
}
