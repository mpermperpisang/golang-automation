package suites

import (
	"github.com/cucumber/godog"
	stepdefinitions "github.com/golang-automation/features/stepdefinitions/apps/ios"
)

func IOSScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user open iOS apps`, stepdefinitions.OpenIOS)
}