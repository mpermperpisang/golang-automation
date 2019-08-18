package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/step_definitions"
	"github.com/golang-automation/features/support"
)

func InitiateAutomation(s *godog.Suite) {
	// web
	s.Step(`^visit dweb$`, step_definitions.OpenDWEB)
	s.Step(`^visit mweb$`, step_definitions.OpenMWEB)
	// apps
	s.Step(`^visit android$`, step_definitions.OpenAndroid)
	s.Step(`^visit ios$`, step_definitions.OpenIOS)
	// api
	s.Step(`^client has "([^\"]*)" as base api$`, step_definitions.BaseAPI)
	s.Step(`^client sends a ([^\"]*) request to "([^\"]*)"$`, step_definitions.RequestAPI)
	s.Step(`^response status should be "([^\"]*)"$`, step_definitions.ResponseStatusCodeAPI)
	// unit
	s.Step(`^user has a name "([^\"]*)"$`, step_definitions.GivenUserName)
	s.Step(`^Testivus meet user$`, step_definitions.MeetUserName)
	s.Step(`^Testivus say "([^\"]*)"$`, step_definitions.SayHelloName)

	support.GodogMainSupport(s)
}
