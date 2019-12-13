package suites

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

/*ExampleAutomation is suites for initial run*/
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
}
