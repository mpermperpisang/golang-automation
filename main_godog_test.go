package main

import (
	"github.com/cucumber/godog"
	apisuites "github.com/golang-automation/features/suites/api"
	androidsuites "github.com/golang-automation/features/suites/apps/android"
	iossuites "github.com/golang-automation/features/suites/apps/ios"
	desktopsuites "github.com/golang-automation/features/suites/web/desktop"
	mobilesuites "github.com/golang-automation/features/suites/web/mobile"
	"github.com/golang-automation/features/supports"
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
