package support

import (
	"fmt"

	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/helper/apps/ios"

	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/helper/web"
)

/*GodogMainSupport does something todo before and after scenario*/
func GodogMainSupport(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		fmt.Println("Preparing scenario")
	})

	s.AfterScenario(func(interface{}, error) {
		if web.Driver != nil {
			web.Driver.Quit()
		} else if android.Driver != nil {
			// android.Device.ClearCookies()
			// android.Device.Reset()
			android.Driver.Stop()
		} else if ios.Driver != nil {
			// ios.Device.ClearCookies()
			// ios.Device.Reset()
			ios.Driver.Stop()
		}

		fmt.Println("Quitting scenario")
	})
}
