package suites

import (
	stepdefinitions "github.com/golang-automation-v1/features/stepdefinitions/apps/ios"

	"github.com/cucumber/godog"
)

func IOSScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user open iOS apps`, stepdefinitions.OpenIOS)
}
