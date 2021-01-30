package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/cucumber/godog"
	"github.com/golang-automation/features/helper"
	apps "github.com/golang-automation/features/helper/apps"
	appsaction "github.com/golang-automation/features/helper/apps/action"
	android "github.com/golang-automation/features/helper/apps/android"
	ios "github.com/golang-automation/features/helper/apps/ios"
	web "github.com/golang-automation/features/helper/web"
	webaction "github.com/golang-automation/features/helper/web/action"
	desktop "github.com/golang-automation/features/helper/web/desktop"
	mobile "github.com/golang-automation/features/helper/web/mobile"
	"github.com/golang-automation/features/support/structs"
)

// DesktopWeb : desktop web page
var DesktopWeb webaction.Page

// MobileWeb : mobile web page
var MobileWeb webaction.Page

// DesktopPage : desktop web page
var DesktopPage desktop.DwebPage

// MobilePage : mobile web page
var MobilePage mobile.MwebPage

// AndroidApps : android apps driver
var AndroidApps appsaction.Page

// IOSApps : ios apps driver
var IOSApps appsaction.Page

// DeviceAction : android & ios action method to device
var DeviceAction appsaction.Page

// Platform : platform name
var Platform string

// PWD : get current directory path
var PWD string
var testCase structs.ScenarioDetail
var path, pathname, filename string
var scenarioTags, featureTags, tags string

func structDetail(scenario interface{}) error {
	data, _ := json.Marshal(scenario)

	return json.Unmarshal(data, &testCase)
}

func ssWeb() error {
	if web.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/web", PWD)

		takeErrorWebPageImage()
	}

	return nil
}

func ssAndroid() error {
	if android.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/android", PWD)

		takeErrorAppsPageImage()
	}

	return nil
}

func ssIOS() error {
	if ios.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/iOS", PWD)

		takeErrorAppsPageImage()
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

func takeErrorWebPageImage() error {
	buffWeb := webaction.Page{Action: web.Driver}

	dirCheck()
	buffWeb.TakeScreenshot()
	ioutil.WriteFile(pathname, []byte(fmt.Sprintf("%v", buffWeb)), 0644)

	return nil
}

func takeErrorAppsPageImage() error {
	buffApps := appsaction.Page{Action: appsaction.AppPage{Page: appsaction.AppsDevice{Device: appsaction.Device}}}

	dirCheck()
	buffApps.TakeScreenshot(pathname)

	return nil
}

// InitializeTestSuite : does something before and after suite
func InitializeTestSuite(s *godog.TestSuiteContext) {
	s.AfterSuite(func() {
		// resp, err := http.Post("http://localhost:8383/godog-support?executionTags="+getExecutionTags()+"&platformName="+getPlatformName()+"&pwdPath="+PWD, "", nil)
		// helper.LogPanicln(err)

		// defer resp.Body.Close()

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

// InitializeScenario : does something before and after scenario
func InitializeScenario(s *godog.ScenarioContext) {
	s.BeforeScenario(func(scenario *godog.Scenario) {
		structDetail(scenario)

		i := 0

		for i < len(testCase.Tags) {
			scenarioTags += " " + testCase.Tags[i].Name

			i++
		}

		platformCheck(scenarioTags)
	})

	s.AfterScenario(func(scenario *godog.Scenario, log error) {
		var err error

		structDetail(scenario)

		scenarioTags = ""
		PWD, err = os.Getwd()
		helper.LogPanicln(err)

		if log != nil {
			filename = fmt.Sprintf("Screenshot - FAILED - %s - %s - %s.png", testCase.Name, testCase.Steps[0].Text, log)

			ssWeb()
			ssAndroid()
			ssIOS()
		}
	})
}

func replacer() *strings.Replacer {
	return strings.NewReplacer(" ", "%20")
}

func getExecutionTags() string {
	return replacer().Replace(fmt.Sprintf("%v", strings.Trim(testCase.Tags[0].Name, " ")))
}

func getPlatformName() string {
	return replacer().Replace(fmt.Sprintf("%v", Platform))
}

func platformCheck(tags string) error {
	if strings.Contains(tags, "@dweb") {
		Platform += "Desktop-Web "

		web.DriverConnect()
		DwebConnect()
	} else if strings.Contains(tags, "@mweb") {
		Platform += "Mobile-Web "

		web.DriverConnect()
		MwebConnect()
	} else if strings.Contains(tags, "@android") {
		Platform += "Android "

		android.DriverConnect()
		AndroidConnect()
		AndroidDevice()
		AppsDevice()
	} else if strings.Contains(tags, "@ios") {
		Platform += "iOS "

		ios.DriverConnect()
		IOSConnect()
		IOSDevice()
		AppsDevice()
	}

	return nil
}

// DwebConnect : connect to dweb driver
func DwebConnect() error {
	DesktopWeb = webaction.Page{Action: web.Driver}
	DesktopPage = desktop.DwebPage{Page: web.Driver}

	return nil
}

// MwebConnect : connect to dweb driver
func MwebConnect() error {
	MobileWeb = webaction.Page{Action: web.Driver}
	MobilePage = mobile.MwebPage{Page: web.Driver}

	return nil
}

// AndroidConnect : connect to android driver
func AndroidConnect() error {
	AndroidApps = appsaction.Page{
		Action: appsaction.AppPage{
			Driver: android.Driver,
		},
	}

	return nil
}

// IOSConnect : connect to iOS driver
func IOSConnect() error {
	IOSApps = appsaction.Page{
		Action: appsaction.AppPage{
			Driver: ios.Driver,
		},
	}

	return nil
}

// AndroidDevice : create new android device
func AndroidDevice() error {
	AndroidApps.StartDriver()
	AndroidApps.NewDevice()

	return nil
}

// IOSDevice : create new iOS device
func IOSDevice() error {
	IOSApps.StartDriver()
	IOSApps.NewDevice()

	return nil
}

// AppsDevice : connect to method to do different action in device
func AppsDevice() error {
	DeviceAction = appsaction.Page{
		Action: appsaction.AppPage{
			Page: appsaction.AppsDevice{
				Device: appsaction.Device,
			},
		},
	}

	return nil
}
