package suites

import (
	"github.com/cucumber/godog"
	stepdefinitions "github.com/mpermperpisang/golang-automation-v1/features/stepdefinitions/apps/ios"
)

func IOSScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user open iOS apps`, stepdefinitions.OpenIOS)
}
