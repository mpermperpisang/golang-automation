package supports

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/cucumber/godog"
	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/data"
	"github.com/golang-automation/features/helper/formats"
	"github.com/golang-automation/features/helper/messages"
	appsbase "github.com/golang-automation/features/supports/base/apps"
	webbase "github.com/golang-automation/features/supports/base/web"
	supports "github.com/golang-automation/features/supports/drivers"
	"github.com/golang-automation/features/supports/logs"
	"github.com/golang-automation/features/supports/screenshots"
	"github.com/golang-automation/features/supports/structs"
)

var (
	testCase               structs.ScenarioDetail
	scenarioTags, platform string
)

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		os.RemoveAll(helper.GetPWD() + formats.TestPath("logs", data.WEB))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("logs", data.APPS))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("screenshots", data.WEB))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("screenshots", data.APPS))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("report", data.WEB))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("report", data.APPS))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("xray", data.WEB))
		os.RemoveAll(helper.GetPWD() + formats.TestPath("xray", data.APPS))
	})

	s.AfterSuite(func() {
		if supports.WebDriver != nil {
			supports.WebDriver.Quit()
		} else if supports.AndroidDriver != nil {
			supports.AndroidDriver.Stop()
		} else if supports.IOSDriver != nil {
			supports.IOSDriver.Stop()
		}
	})
}

func InitializeScenario(s *godog.ScenarioContext) {
	s.BeforeScenario(func(scenario *godog.Scenario) {
		scenarioDetail(scenario)

		for i := 0; i < len(testCase.Tags); i++ {
			scenarioTags += " " + testCase.Tags[i].Name
		}

		platformCheck(scenarioTags)
	})

	s.AfterScenario(func(scenario *godog.Scenario, log error) {
		scenarioTags = ""

		scenarioDetail(scenario)

		if log != nil {
			createSS()
			createLog(log)
		}
	})
}

func scenarioDetail(scenario interface{}) error {
	data, err := json.Marshal(scenario)
	helper.LogPanicln(err)

	return json.Unmarshal(data, &testCase)
}

func platformCheck(tags string) error {
	if strings.Contains(tags, "@"+data.API) {
		platform = data.API
	} else if strings.Contains(tags, "@"+data.DWEB) {
		platform = data.DWEB

		webStart(platform)
	} else if strings.Contains(tags, "@"+data.MWEB) {
		platform = data.MWEB

		webStart(platform)
	} else if strings.Contains(tags, "@"+data.ANDROID) {
		platform = data.ANDROID

		appsStart(platform)
	} else if strings.Contains(tags, "@"+data.IOS) {
		platform = data.IOS

		appsStart(platform)
	} else {
		helper.LogPanicln(messages.PlatformList())
	}

	return nil
}

func webStart(platform string) {
	supports.WebCapabilities()
	webbase.OpenWebURL(platform, "/")
}

func appsStart(platform string) {
	switch platform {
	case data.ANDROID:
		supports.AndroidCapabilities()
	case data.IOS:
		supports.IOSCapabilities()
	default:
		helper.LogPanicln(messages.PlatformList())
	}

	appsbase.OpenApps(platform)
}

func createSS() {
	helper.SetFilename("ss", platform, testCase.Name)
	screenshots.SSWeb(platform)
	screenshots.SSAndroid()
	screenshots.SSIOS()
}

func createLog(log error) {
	helper.SetFilename("log", platform, testCase.Name)
	logs.LogWeb(platform, log)
	logs.LogAndroid(log)
	logs.LogIOS(log)
}
