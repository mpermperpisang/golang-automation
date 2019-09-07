package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/suites"
	"github.com/golang-automation/features/support"
)

func GodogMainSuites(s *godog.Suite) {
	suites.InitiateAutomation(s)
	suites.AutomationGlobal(s)
	suites.AutomationAPI(s)
	suites.AutomationWeb(s)
	suites.AutomationApps(s)

	support.GodogMainSupport(s)
}
