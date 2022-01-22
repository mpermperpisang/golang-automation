package drivers

import (
	"fmt"

	"github.com/magiconair/properties"
	"github.com/tebeka/selenium"

	"github.com/golang-automation-v1/features/helper"
)

var WebDriver selenium.WebDriver

func webProperties() *properties.Properties {
	return properties.MustLoadFile(helper.GetPWD()+"/capabilities-web.properties", properties.UTF8)
}

func WebCapabilities() {
	var err error

	fmt.Println("1")
	p := webProperties()
	fmt.Println("2")

	caps := selenium.Capabilities{
		"browserName": p.MustGetString("browserName"),
	}
	fmt.Println("3")

	WebDriver, err = selenium.NewRemote(caps, "")
	fmt.Println(err)
	helper.LogPanicln(err)
}
