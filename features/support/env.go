package support

import (
	"fmt"

	"github.com/DATA-DOG/godog"
	android "github.com/golang-automation/features/helper/apps/android"
	ios "github.com/golang-automation/features/helper/apps/ios"
	web "github.com/golang-automation/features/helper/web"
)

/*GodogMainSupport does something todo before and after scenario*/
func GodogMainSupport(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		fmt.Println("Starting scenario")
	})

	s.AfterScenario(func(interface{}, error) {
		if web.Driver != nil {
			web.Driver.Screenshot()
			web.Driver.Quit()
		} else if android.Driver != nil {
			android.Device.Screenshot("error_android.jpg")
			android.Driver.Stop()
		} else if ios.Driver != nil {
			ios.Device.Screenshot("error_ios.jpg")
			ios.Driver.Stop()
		}

		fmt.Println("Finishing scenario")
	})
}
