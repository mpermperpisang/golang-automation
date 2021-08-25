package suites

import (
	stepdefinitions "golang-automation/features/stepdefinitions/apps/android"

	"github.com/cucumber/godog"
)

func AndroidScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user click menu`, stepdefinitions.ClickMenu)
	s.Step(`validate action bar tabs is displaying`, stepdefinitions.ValidateActionBarTabs)
}
