package appshelper

import (
	"os"
	"os/exec"

	"github.com/golang-API/helper"
)

/*Appium : set appium server*/
func Appium() error {
	pwd, err := os.Getwd()
	helper.LogPanicln(err)

	killADB := exec.Command("/bin/sh", pwd+"/script/appium.sh").Run()
	helper.LogPanicln(killADB)

	return nil
}
