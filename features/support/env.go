package support

import (
	"fmt"
	"os"
	"regexp"

	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/helper"
	android "github.com/golang-automation/features/helper/apps/android"
	androidDriver "github.com/golang-automation/features/helper/apps/android/driver"
	ios "github.com/golang-automation/features/helper/apps/ios"
	iosDriver "github.com/golang-automation/features/helper/apps/ios/driver"
	web "github.com/golang-automation/features/helper/web"
)

/*GodogMainSupport does something todo before and after scenario*/
func GodogMainSupport(s *godog.Suite) {
	// TODO: Handle before and after scenario
	s.BeforeScenario(func(interface{}) {
		argsWithProg := os.Args
		tag := regexp.MustCompile(helper.RegexTag()).FindString(fmt.Sprint(argsWithProg))

		fmt.Println("Starting automation")
		fmt.Println("Running scenario with tag : " + tag)
	})
	

	s.AfterScenario(func(interface{}, error) {
		if web.Driver != nil {
			web.Driver.Screenshot()
			web.Driver.Quit()
		} else if android.Driver != nil {
			androidDriver.Device.Screenshot("error_android.jpg")
			android.Driver.Stop()
		} else if ios.Driver != nil {
			iosDriver.Device.Screenshot("error_ios.jpg")
			ios.Driver.Stop()
		}

		fmt.Println("Stopping automation")
	})
}
