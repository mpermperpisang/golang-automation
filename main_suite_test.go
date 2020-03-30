package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/suites"
	android "github.com/golang-automation/features/suites/android"
	api "github.com/golang-automation/features/suites/api"
	desktop "github.com/golang-automation/features/suites/desktop"
	mobile "github.com/golang-automation/features/suites/mobile"
	"github.com/golang-automation/features/support"
)

func GodogMainSuites(s *godog.Suite) {
	// general suites
	suites.ExampleAutomation(s)
	suites.AutomationGlobal(s)
	suites.FarewellGreeting(s)
	// suites by platform
	api.AutomationAPI(s)
	desktop.AutomationDesktop(s)
	mobile.AutomationMobile(s)
	android.AutomationAndroid(s)

	support.GodogMainSupport(s)
}


