package appshelper

import "os/exec"

/*AppiumStop : stop appium server*/
func AppiumStop() error {
	exec.Command("/bin/sh", "../script/appium.sh")

	return nil
}
