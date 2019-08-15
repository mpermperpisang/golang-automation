package main

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/helper/web"
)

func GodogMainSupport(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		fmt.Println("Preparing for run automation")
	})

	s.AfterScenario(func(interface{}, error) {
		fmt.Println("Quitting driver")
		if web.Driver != nil {
			web.Driver.Quit()
		}
	})
}
