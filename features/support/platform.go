package support

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/gherkin"
	apps "github.com/golang-automation/features/helper/apps"
	"github.com/golang-automation/features/helper/apps/action"
	android "github.com/golang-automation/features/helper/apps/android"
	ios "github.com/golang-automation/features/helper/apps/ios"
	"github.com/golang-automation/features/helper/errors"
	web "github.com/golang-automation/features/helper/web"
	desktop "github.com/golang-automation/features/helper/web/desktop"
	mobile "github.com/golang-automation/features/helper/web/mobile"
)

// Dweb : desktop web driver
var Dweb web.WebDriver

// Mweb : mobile web driver
var Mweb web.WebDriver

// AndroidPage : android apps page
var AndroidPage action.AppPage

// IOSPage : ios apps page
var IOSPage action.AppPage

// DesktopPage : desktop web page
var DesktopPage desktop.DwebPage

// MobilePage : mobile web page
var MobilePage mobile.MwebPage

var androidApps, iOSApps action.Driver
var testCase scenarioDetail
var feature featureDetail
var path, pathname, pwd, filename string
var tags string

func structDetail(scenario interface{}, typeSupport string) error {
	data, _ := json.Marshal(scenario)

	if typeSupport == "tc" {
		json.Unmarshal(data, &testCase)
	} else {
		json.Unmarshal(data, &feature)
	}

	return nil
}

func ssWeb() error {
	if web.Driver != nil {
		path = fmt.Sprintf("%s/screenshots/web", pwd)
		buffWeb, err := web.Driver.Screenshot()
		errors.LogPanicln(err)

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

// GodogMainSupport : does something before and after scenario
func GodogMainSupport(s *godog.Suite) {
	// TODO: Handle before and after scenario
	s.BeforeFeature(func(scenario *gherkin.Feature) {
		structDetail(scenario, "feature")

		i := 0

		for i < len(feature.Tags) {
			tags += " " + feature.Tags[i].Name

			i++
		}
	})

	s.BeforeScenario(func(scenario interface{}) {
		structDetail(scenario, "tc")

		i := 0

		for i < len(testCase.Tags) {
			tags += " " + testCase.Tags[i].Name

			i++
		}

		if tags != "" {
			fmt.Println("Starting automation")
			fmt.Printf("Running scenario with tag :%s", tags)
		}

		platformCheck(tags)
	})

	s.AfterScenario(func(scenario interface{}, log error) {
		var err error

		structDetail(scenario, "tc")

		if log != nil {
			pwd, err = os.Getwd()
			errors.LogPanicln(err)

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

func platformCheck(tags string) error {
	if strings.Contains(tags, "@dweb") {
		DwebConnect()
	} else if strings.Contains(tags, "@mweb") {
		MwebConnect()
	} else if strings.Contains(tags, "@android") {
		AndroidConnect()
	} else if strings.Contains(tags, "@ios") {
		IOSConnect()
	}

	return nil
}

// DwebConnect : connect to dweb driver
func DwebConnect() error {
	driver := web.DriverConnect()
	Dweb = web.WebDriver{Driver: driver}
	DesktopPage = desktop.DwebPage{Page: Dweb}

	return nil
}

// MwebConnect : connect to dweb driver
func MwebConnect() error {
	driver := web.DriverConnect()
	Mweb = web.WebDriver{Driver: driver}
	MobilePage = mobile.MwebPage{Page: Mweb}

	return nil
}

// AndroidConnect : connect to androids driver
func AndroidConnect() error {
	driver := android.DriverConnect()
	androidApps = action.Driver{Driver: driver}
	AndroidPage = action.AppPage{Page: androidApps}

	return nil
}

// IOSConnect : connect to iOS driver
func IOSConnect() error {
	driver := ios.DriverConnect()
	iOSApps = action.Driver{Driver: driver}
	IOSPage = action.AppPage{Page: iOSApps}

	return nil
}

// AndroidOpenDevice : open android device
func AndroidOpenDevice() error {
	AndroidPage.StartDriver()
	AndroidPage.NewDevice()

	return nil
}

// IOSOpenDevice : open iOS device
func IOSOpenDevice() error {
	IOSPage.StartDriver()
	IOSPage.NewDevice()

	return nil
}
