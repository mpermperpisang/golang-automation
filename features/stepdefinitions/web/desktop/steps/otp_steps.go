package step

import (
	"os"
	"time"

	web "github.com/golang-automation/features/helper/web/action"
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
)

/*OTPConfirm function to input OTP if exist*/
func OTPConfirm() error {
	web.FindElementByClickScript(desktoppages.BtnSend)
	time.Sleep(time.Second * 1)
	web.SendKeysByID(desktoppages.FieldOTP, os.Getenv("OTP"))
	web.FindElementByClickScript(desktoppages.BtnConfirm)
	web.FindElementByClickScript(desktoppages.BtnOke)

	return nil
}
