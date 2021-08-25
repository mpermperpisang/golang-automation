package main

import (
	"github.com/cucumber/godog"
	apisuites "github.com/mpermperpisang/golang-automation-v1/features/suites/api"
	androidsuites "github.com/mpermperpisang/golang-automation-v1/features/suites/apps/android"
	iossuites "github.com/mpermperpisang/golang-automation-v1/features/suites/apps/ios"
	desktopsuites "github.com/mpermperpisang/golang-automation-v1/features/suites/web/desktop"
	mobilesuites "github.com/mpermperpisang/golang-automation-v1/features/suites/web/mobile"
	"github.com/mpermperpisang/golang-automation-v1/features/supports"
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
