package suites

import (
	"github.com/cucumber/godog"
	stepdefinitions "github.com/mpermperpisang/golang-automation-v1/features/stepdefinitions/web/desktop"
)

func DwebScenarioContext(s *godog.ScenarioContext) {
	s.Step(`user visit "([^\"]*)" link`, stepdefinitions.VisitLink)
	s.Step(`user click "([^\"]*)" button`, stepdefinitions.ClickButton)
	s.Step(`validate login warning is displaying`, stepdefinitions.ValidateLoginWarning)
}
