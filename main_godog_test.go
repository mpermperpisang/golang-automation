package main

import (
	apisuites "golang-automation/features/suites/api"
	androidsuites "golang-automation/features/suites/apps/android"
	iossuites "golang-automation/features/suites/apps/ios"
	desktopsuites "golang-automation/features/suites/web/desktop"
	mobilesuites "golang-automation/features/suites/web/mobile"
	"golang-automation/features/supports"

	"github.com/cucumber/godog"
)

func MainTestSuiteContext(s *godog.TestSuiteContext) {
	supports.InitializeTestSuite(s)
}

func MainScenarioContext(s *godog.ScenarioContext) {
	supports.InitializeScenario(s)

	apisuites.APIScenarioContext(s)
	desktopsuites.DwebScenarioContext(s)
	mobilesuites.MwebScenarioContext(s)
	androidsuites.AndroidScenarioContext(s)
	iossuites.IOSScenarioContext(s)
}
