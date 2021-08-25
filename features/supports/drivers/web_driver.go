package drivers

import (
	"golang-automation/features/helper"

	"github.com/magiconair/properties"
	"github.com/tebeka/selenium"
)

var WebDriver selenium.WebDriver

func webProperties() *properties.Properties {
	return properties.MustLoadFile(helper.GetPWD()+"/capabilities-web.properties", properties.UTF8)
}

func WebCapabilities() {
	var err error

	p := webProperties()

	caps := selenium.Capabilities{
		"browserName": p.MustGetString("browserName"),
	}

	WebDriver, err = selenium.NewRemote(caps, "")
	helper.LogPanicln(err)
}
