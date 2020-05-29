package desktoppages

import (
	"os"
	"time"

	webaction "github.com/golang-automation/features/helper/web/action"
)

/*OTPPage : page object one time password*/
type OTPPage struct {
	Page webaction.Page
}

var (
	fieldOTP       = "otp"
	labelInputCode = "Masukkan kode ke kotak isian di bawah ini "
	btnSend        = ".c-btn__text"
	btnConfirm     = ".js-tfa-submit"
	btnOke         = ".c-btn--red.js-tfa-step--finish"
)

/*ClickSend : click send otp button*/
func (s *OTPPage) ClickSend() *OTPPage {
	s.Page.FindElementByClickScript(btnSend)

	return &OTPPage{Page: s.Page}
}

/*InputOTP : input one time password code*/
func (s *OTPPage) InputOTP() *OTPPage {
	time.Sleep(time.Second * 1)
	s.Page.SendKeysByID(fieldOTP, os.Getenv("OTP"))

	return &OTPPage{Page: s.Page}
}

/*ClickConfirm : click confirmation button*/
func (s *OTPPage) ClickConfirm() *OTPPage {
	s.Page.FindElementByClickScript(btnConfirm)

	return &OTPPage{Page: s.Page}
}

/*ClickOK : click oke button*/
func (s *OTPPage) ClickOK() *HomePage {
	s.Page.FindElementByClickScript(btnOke)

	return &HomePage{Page: s.Page}
}
