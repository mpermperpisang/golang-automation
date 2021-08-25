package suites

import (
	stepdefinitions "golang-automation/features/stepdefinitions/api"

	"github.com/cucumber/godog"
)

func APIScenarioContext(s *godog.ScenarioContext) {
	s.Step(`client sends a ([^\"]*) to "([^\"]*)"`, stepdefinitions.HitEndpointWithoutBody)
	s.Step(`client sends a ([^\"]*) to "([^\"]*)" with body:`, stepdefinitions.HitEndpointWithBody)
	s.Step(`validate status is (\d+)`, stepdefinitions.ValidateStatus)
	s.Step(`validate "([^\"]*)" is "([^\"]*)"`, stepdefinitions.ValidateValue)
}
