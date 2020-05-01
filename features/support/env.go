package support

import (
	"fmt"
	"os"
	"regexp"

	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/helper"
	appshelper "github.com/golang-automation/features/helper/apps"
	android "github.com/golang-automation/features/helper/apps/android"
	ios "github.com/golang-automation/features/helper/apps/ios"
	web "github.com/golang-automation/features/helper/web"
)

/*GodogMainSupport does something todo before and after scenario*/
func GodogMainSupport(s *godog.Suite) {
	// TODO: Handle before and after scenario
	s.BeforeScenario(func(interface{}) {
		argsWithProg := os.Args
		tag := regexp.MustCompile(helper.RegexTag()).FindString(fmt.Sprint(argsWithProg))

		fmt.Println("Starting automation")
		if tag != "" {
			fmt.Println("Running scenario with tag : " + tag)
		}
	})

	s.AfterScenario(func(interface{}, error) {
		fmt.Println("Stopping automation")

		if web.Driver != nil {
			web.Driver.Quit()
		} else if android.Driver != nil {
			android.Driver.Stop()
			appshelper.AppiumStop()
		} else if ios.Driver != nil {
			ios.Driver.Stop()
			appshelper.AppiumStop()
		}
	})
}
