package mobilepages

import (
	"os"
	"time"

	webaction "github.com/golang-automation/features/helper/web/action"
)

// OTPPage : page object otp
type OTPPage struct {
	Page webaction.Page
}

var (
	fieldOTP   = "otp"
	btnSend    = ".js-tfa-request__button.js-tfa-step.js-tfa-step--request"
	btnLanjut  = ".js-tfa-submit"
	btnLanjut2 = ".js-tfa-step--finish"
)

// ClickSend : click send otp button
func (s OTPPage) ClickSend() *OTPPage {
	time.Sleep(time.Second * 1)
	s.Page.FindElementByClickScript(btnSend)

	return &OTPPage{Page: s.Page}
}

// InputOTP : input one time password code
func (s OTPPage) InputOTP() *OTPPage {
	time.Sleep(time.Second * 1)
	s.Page.SendKeysByID(fieldOTP, os.Getenv("OTP"))

	return &OTPPage{Page: s.Page}
}

// ClickNext : click Lanjut button
func (s OTPPage) ClickNext() *HomePage {
	s.Page.FindElementByClickScript(btnLanjut)
	time.Sleep(time.Second * 1)
	s.Page.FindElementByClickScript(btnLanjut2)

	return &HomePage{Page: s.Page}
}
