package main

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/step_definitions"
)

func VariantAPI(s *godog.Suite) {
	s.Step(`^response should have "([^"]*)"$`, step_definitions.ResponseMatchingKey)
	s.Step(`^response should have "([^"]*)" matching "([^"]*)"$`, step_definitions.ResponseMatchingValue)
}
