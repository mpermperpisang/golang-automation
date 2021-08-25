package suites

import (
	stepdefinitions "golang-automation/features/stepdefinitions/web/desktop"

	"github.com/cucumber/godog"
)

func DwebScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user visit "([^\"]*)" link`, stepdefinitions.VisitLink)
	s.Step(`user click "([^\"]*)" button`, stepdefinitions.ClickButton)
	s.Step(`validate login warning is displaying`, stepdefinitions.ValidateLoginWarning)
}
