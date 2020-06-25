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
	"github.com/golang-automation/features/helper/apps/action"
	android "github.com/golang-automation/features/helper/apps/android"
	ios "github.com/golang-automation/features/helper/apps/ios"
	web "github.com/golang-automation/features/helper/web"
	"github.com/logrusorgru/aurora"
)

var testCase testCaseDetail
var path, pathname, pwd, filename string

func scenarioDetail(scenario interface{}) error {
	data, _ := json.Marshal(scenario)

	json.Unmarshal(data, &testCase)

	return nil
}

func ssWeb() error {
	if web.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/web", pwd)
		buffWeb, err := web.Driver.Screenshot()
		helper.LogPanicln(err)

		dirCheck()
		ioutil.WriteFile(pathname, buffWeb, 0644)
	}

	return nil
}

func ssAndroid() error {
	if android.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/android", pwd)

		takeErrorPageImage()
	}

	return nil
}

func ssIOS() error {
	if ios.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/iOS", pwd)

		takeErrorPageImage()
	}

	return nil
}

func dirCheck() error {
	pathname = filepath.Join(path, filename)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}

	return nil
}

func takeErrorPageImage() error {
	buffApps := action.Page{Action: action.AppPage{Device: action.Device}}

	dirCheck()
	buffApps.TakeScreenshot(pathname)

	return nil
}

/*GodogMainSupport : does something before and after scenario*/
func GodogMainSupport(s *godog.Suite) {
	// TODO: Handle before and after scenario
	s.BeforeScenario(func(scenario interface{}) {
		var tags string

		scenarioDetail(scenario)

		i := 0
		name := aurora.Italic(aurora.Bold(aurora.White(testCase.Name)))

		for i < len(testCase.Tags) {
			tags += " " + testCase.Tags[i].Name

			i++
		}

		if tags != "" {
			fmt.Println("Starting automation")
			fmt.Printf("Running %s scenario :%s", name, tags)
		}
	})

	s.AfterScenario(func(scenario interface{}, log error) {
		var err error

		scenarioDetail(scenario)

		if log != nil {
			pwd, err = os.Getwd()
			helper.LogPanicln(err)

			filename = fmt.Sprintf("FAILED - %s - %s - %s.png", testCase.Name, testCase.Steps[0].Text, log)

			ssWeb()
			ssAndroid()
			ssIOS()
		}
	})

	s.AfterSuite(func() {
		fmt.Println("Stopping automation")

		if web.Driver != nil {
			web.Driver.Quit()
		} else if android.Driver != nil {
			android.Driver.Stop()
			apps.Appium()
		} else if ios.Driver != nil {
			ios.Driver.Stop()
			apps.Appium()
		}
	})
}
