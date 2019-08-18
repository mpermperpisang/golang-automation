package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/suites"
)

func GodogMainSuites(s *godog.Suite) {
	suites.InitiateAutomation(s)
	suites.AutomationAPI(s)
}
