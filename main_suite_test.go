package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/suites"
	"github.com/golang-automation/features/suites/androidsuites"
	"github.com/golang-automation/features/suites/apisuites"
	"github.com/golang-automation/features/suites/desktopsuites"
	"github.com/golang-automation/features/suites/mobilesuites"
	"github.com/golang-automation/features/support"
)

func GodogMainSuites(s *godog.Suite) {
	// generic suites
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
