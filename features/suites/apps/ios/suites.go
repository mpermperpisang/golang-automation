package suites

import (
	stepdefinitions "golang-automation/features/stepdefinitions/apps/ios"

	"github.com/cucumber/godog"
)

func IOSScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user open iOS apps`, stepdefinitions.OpenIOS)
}
