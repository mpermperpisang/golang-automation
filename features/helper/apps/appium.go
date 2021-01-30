package appshelper

import (
	"os"
	"os/exec"

	"github.com/golang-automation/features/helper"
)

// Appium : set appium server
func Appium() error {
	pwd, err := os.Getwd()
	helper.LogPanicln(err)

	return exec.Command("/bin/sh", pwd+"/features/support/scripts/appium.sh").Run()
}
