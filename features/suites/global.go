package suites

import (
	"github.com/cucumber/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

// AutomationGlobal : suites for dweb/mweb/android/ios
func AutomationGlobal(s *godog.ScenarioContext) {
	s.Step(`^there is client who wants to login as "([^"]*)" via (desktop|mobile|android|ios)$`, stepdefinitions.LoginData)
}
