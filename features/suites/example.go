package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

// ExampleAutomation : suites for initial run
func ExampleAutomation(s *godog.Suite) {
	// web
	s.Step(`^visit dweb$`, stepdefinitions.OpenDWEB)
	s.Step(`^visit mweb$`, stepdefinitions.OpenMWEB)
	// apps
	s.Step(`^visit android$`, stepdefinitions.OpenAndroid)
	s.Step(`^visit ios$`, stepdefinitions.OpenIOS)
	// api
	s.Step(`^client has "([^\"]*)" as base api$`, stepdefinitions.BaseAPI)
	s.Step(`^response status should be "([^\"]*)"$`, stepdefinitions.ResponseStatusAPI)
	// unit
	s.Step(`^user has a name "([^\"]*)"$`, stepdefinitions.GivenUserName)
	s.Step(`^Testivus meet user$`, stepdefinitions.MeetUserName)
	s.Step(`^Testivus say "([^\"]*)"$`, stepdefinitions.SayHelloName)
	// data
	s.Step(`^some data$`, stepdefinitions.GivenData)
	s.Step(`^set data$`, stepdefinitions.SetData)
	s.Step(`^get data$`, stepdefinitions.GetData)
	// yaml
	s.Step(`^yaml file$`, stepdefinitions.GivenFile)
	s.Step(`^read file$`, stepdefinitions.ReadFile)
	s.Step(`^print contents$`, stepdefinitions.PrintContents)
	// mapping
	s.Step(`^function mapping$`, stepdefinitions.GetFunction)
}
