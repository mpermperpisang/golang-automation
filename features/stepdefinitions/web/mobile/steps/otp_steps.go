package step

import (
	"os"
	"time"

	web "github.com/golang-automation/features/helper/web/action"
	mobilepages "github.com/golang-automation/features/objectabstractions/web/mobile"
)

/*OTPConfirm function to input OTP if exist*/
func OTPConfirm() error {
	time.Sleep(time.Second * 1)
	web.FindElementByClickScript(mobilepages.BtnSend)
	time.Sleep(time.Second * 1)
	web.SendKeysByID(mobilepages.FieldOTP, os.Getenv("OTP"))
	web.FindElementByClickScript(mobilepages.BtnLanjut)
	time.Sleep(time.Second * 1)
	web.FindElementByClickScript(mobilepages.BtnLanjut2)

	return nil
}
