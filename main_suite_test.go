package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/suites"
	androidsuites "github.com/golang-automation/features/suites/android"
	apisuites "github.com/golang-automation/features/suites/api"
	desktopsuites "github.com/golang-automation/features/suites/desktop"
	mobilesuites "github.com/golang-automation/features/suites/mobile"
	"github.com/golang-automation/features/support"
)

func GodogMainSuites(s *godog.Suite) {
	// general suites
	suites.ExampleAutomation(s)
	suites.AutomationGlobal(s)
	suites.FarewellGreeting(s)
	// suites by platform
	apisuites.AutomationAPI(s)
	desktopsuites.AutomationDesktop(s)
	mobilesuites.AutomationMobile(s)
	androidsuites.AutomationAndroid(s)

	support.GodogMainSupport(s)
}
