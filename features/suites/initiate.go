package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
	"github.com/golang-automation/features/support"
)

func InitiateAutomation(s *godog.Suite) {
	// web
	s.Step(`^visit dweb$`, stepdefinitions.OpenDWEB)
	s.Step(`^visit mweb$`, stepdefinitions.OpenMWEB)
	// apps
	s.Step(`^visit android$`, stepdefinitions.OpenAndroid)
	s.Step(`^visit ios$`, stepdefinitions.OpenIOS)
	// api
	s.Step(`^client has "([^\"]*)" as base api$`, stepdefinitions.BaseAPI)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, stepdefinitions.RequestAPI)
	s.Step(`^response status should be "([^\"]*)"$`, stepdefinitions.ResponseStatusCodeAPI)
	// unit
	s.Step(`^user has a name "([^\"]*)"$`, stepdefinitions.GivenUserName)
	s.Step(`^Testivus meet user$`, stepdefinitions.MeetUserName)
	s.Step(`^Testivus say "([^\"]*)"$`, stepdefinitions.SayHelloName)

	support.GodogMainSupport(s)
}
