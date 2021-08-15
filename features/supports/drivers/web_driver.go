package drivers

import (
	"github.com/golang-automation/features/helper"
	"github.com/magiconair/properties"
	"github.com/tebeka/selenium"
)

var WebDriver selenium.WebDriver

func WebCapabilities() {
	var err error

	p := properties.MustLoadFile("${GOPATH}/src/github.com/golang-automation/capabilities-web.properties", properties.UTF8)
	caps := selenium.Capabilities{"browserName": p.MustGetString("browserName")}
	WebDriver, err = selenium.NewRemote(caps, "")
	helper.LogPanicln(err)
}
