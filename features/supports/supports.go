package supports

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/cucumber/godog"
	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/formats"
	androidbase "github.com/golang-automation/features/supports/base/apps/android"
	iosbase "github.com/golang-automation/features/supports/base/apps/ios"
	webbase "github.com/golang-automation/features/supports/base/web"
	supports "github.com/golang-automation/features/supports/drivers"
	"github.com/golang-automation/features/supports/logs"
	"github.com/golang-automation/features/supports/screenshots"
	"github.com/golang-automation/features/supports/structs"
)

var (
	testCase               structs.ScenarioDetail
	scenarioTags, platform string
	Filename, prefix       string
)

func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.BeforeSuite(func() {
		os.RemoveAll(helper.GetPWD() + formats.LogPath("web"))
		os.RemoveAll(helper.GetPWD() + formats.LogPath("apps"))
		os.RemoveAll(helper.GetPWD() + formats.SSPath("web"))
		os.RemoveAll(helper.GetPWD() + formats.SSPath("apps"))
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
	if strings.Contains(tags, "@dweb") {
		platform += "Desktop-Web "
		prefix = "DESKTOP"

		webStart("dweb")
	} else if strings.Contains(tags, "@mweb") {
		platform += "Mobile-Web "
		prefix = "MOBILE"

		webStart("mweb")
	} else if strings.Contains(tags, "@android") {
		platform += "Android "
		prefix = "ANDROID"

		androidStart()
	} else if strings.Contains(tags, "@ios") {
		platform += "iOS "
		prefix = "IOS"

		iosStart()
	}

	return nil
}

func webStart(platform string) {
	supports.WebCapabilities()
	webbase.OpenWebURL(platform, "/")
}

func androidStart() {
	supports.AndroidCapabilities()
	androidbase.OpenAndroidApps()
}

func iosStart() {
	supports.IOSCapabilities()
	iosbase.OpenIOSApps()
}

func createSS() {
	helper.SetFilename("ss", prefix, testCase.Name)
	screenshots.SSWeb(prefix)
	screenshots.SSAndroid()
	screenshots.SSIOS()
}

func createLog(log error) {
	helper.SetFilename("log", prefix, testCase.Name)
	logs.LogWeb(prefix, log)
	logs.LogAndroid(log)
	logs.LogIOS(log)
}
