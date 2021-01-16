package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

// AppsDevice : ios apps device
var AppsDevice appsaction.Page

// Platform : platform name
var Platform string

// PWD : get current directory path
var PWD string
var testCase scenarioDetail
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
	buffApps := appsaction.Page{Action: appsaction.AppPage{Device: appsaction.Device}}

	dirCheck()
	buffApps.TakeScreenshot(pathname)

	return nil
}

// GodogMainSupport : does something before and after scenario
func GodogMainSupport(s *godog.Suite) {
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

	s.AfterSuite(func() {
		resp, err := http.Post("http://localhost:8383/godog-support?executionTags="+getExecutionTags()+"&platformName="+getPlatformName()+"&pwdPath="+PWD, "", nil)
		helper.LogPanicln(err)

		defer resp.Body.Close()

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

		DwebConnect()
	} else if strings.Contains(tags, "@mweb") {
		Platform += "Mobile-Web "

		MwebConnect()
	} else if strings.Contains(tags, "@android") {
		Platform += "Android "

		AndroidConnect()
		AndroidOpenDevice()
		appsDevice()
	} else if strings.Contains(tags, "@ios") {
		Platform += "iOS "

		IOSConnect()
		IOSOpenDevice()
		appsDevice()
	}

	return nil
}

// DwebConnect : connect to dweb driver
func DwebConnect() error {
	web.DriverConnect()

	DesktopWeb = webaction.Page{Action: web.Driver}
	DesktopPage = desktop.DwebPage{Page: web.Driver}

	return nil
}

// MwebConnect : connect to dweb driver
func MwebConnect() error {
	web.DriverConnect()

	MobileWeb = webaction.Page{Action: web.Driver}
	MobilePage = mobile.MwebPage{Page: web.Driver}

	return nil
}

// AndroidConnect : connect to androids driver
func AndroidConnect() error {
	android.DriverConnect()

	AndroidApps = appsaction.Page{
		Action: appsaction.AppPage{
			Page: appsaction.Driver{
				Driver: android.Driver,
			},
		},
	}
	return nil
}

// IOSConnect : connect to iOS driver
func IOSConnect() error {
	ios.DriverConnect()
	IOSOpenDevice()

	IOSApps = appsaction.Page{
		Action: appsaction.AppPage{
			Page: appsaction.Driver{
				Driver: ios.Driver,
			},
		},
	}

	return nil
}

// AndroidOpenDevice : open android device
func AndroidOpenDevice() error {
	AndroidApps.StartDriver()
	AndroidApps.NewDevice()

	return nil
}

// IOSOpenDevice : open iOS device
func IOSOpenDevice() error {
	IOSApps.StartDriver()
	IOSApps.NewDevice()

	return nil
}

func appsDevice() error {
	AppsDevice = appsaction.Page{
		Action: appsaction.AppPage{
			Device: appsaction.Device,
		},
	}

	return nil
}
