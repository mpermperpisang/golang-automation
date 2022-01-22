package main

import (
	apisuites "github.com/golang-automation-v1/features/suites/api"
	androidsuites "github.com/golang-automation-v1/features/suites/apps/android"
	iossuites "github.com/golang-automation-v1/features/suites/apps/ios"
	desktopsuites "github.com/golang-automation-v1/features/suites/web/desktop"
	mobilesuites "github.com/golang-automation-v1/features/suites/web/mobile"
	"github.com/golang-automation-v1/features/supports"

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
