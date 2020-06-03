package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/helper"
	apps "github.com/golang-automation/features/helper/apps"
	android "github.com/golang-automation/features/helper/apps/android"
	ios "github.com/golang-automation/features/helper/apps/ios"
	web "github.com/golang-automation/features/helper/web"
	"github.com/logrusorgru/aurora"
)

var testCase testCaseDetail

func scenarioDetail(scenario interface{}) error {
	data, _ := json.Marshal(scenario)

	json.Unmarshal(data, &testCase)

	return nil
}

/*GodogMainSupport : does something before and after scenario*/
func GodogMainSupport(s *godog.Suite) {
	// TODO: Handle before and after scenario
	s.BeforeScenario(func(scenario interface{}) {
		scenarioDetail(scenario)

		name := aurora.Italic(aurora.Bold(aurora.White(testCase.Name)))
		tags := testCase.Tags[0].Name

		if tags != "" {
			fmt.Println("Starting automation")
			fmt.Printf("Running scenario %s : %s", name, tags)
		}
	})

	s.AfterScenario(func(scenario interface{}, log error) {
		scenarioDetail(scenario)

		if log != nil {
			buff, err := web.Driver.Screenshot()
			helper.LogPanicln(err)

			pwd, err := os.Getwd()
			helper.LogPanicln(err)

			path := fmt.Sprintf("%s/screenshots", pwd)
			filename := fmt.Sprintf("FAILED - %s - %s - %s.png", testCase.Name, testCase.Steps[0].Text, log)
			pathname := filepath.Join(path, filename)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				os.MkdirAll(path, 0755)
			}

			ioutil.WriteFile(pathname, buff, 0644)
		}
	})

	s.AfterSuite(func() {
		fmt.Println("Stopping automation")

		if web.Driver != nil {
			web.Driver.Quit()
		} else if android.Driver != nil {
			android.Driver.Stop()
			apps.AppiumStop()
		} else if ios.Driver != nil {
			ios.Driver.Stop()
			apps.AppiumStop()
		}
	})
}
