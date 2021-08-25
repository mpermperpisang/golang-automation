package suites

import (
	"github.com/cucumber/godog"
	stepdefinitions "github.com/golang-automation/features/stepdefinitions/apps/android"
)

func AndroidScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user click menu`, stepdefinitions.ClickMenu)
	s.Step(`validate action bar tabs is displaying`, stepdefinitions.ValidateActionBarTabs)
}