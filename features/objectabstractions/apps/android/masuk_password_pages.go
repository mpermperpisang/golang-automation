package androidpages

import (
	"os"
	"time"

	android "github.com/golang-automation/features/helper/apps/android/action"
)

/*InputPasswordPage : page object input password*/
type InputPasswordPage struct {
	Page android.Page
}

var (
	fieldPassword    = "//*[@resource-id='app:id/textFieldAV_textField']"
	btnMasukPassword = "Masuk"
)

/*InputPassword : input password code*/
func (s *InputPasswordPage) InputPassword() *InputPasswordPage {
	s.Page.SendKeysByXpath(fieldPassword, os.Getenv("USER_PHONE_PASSWORD"))

	return &InputPasswordPage{Page: s.Page}
}

/*ClickMasuk : click enter to login*/
func (s *InputPasswordPage) ClickMasuk() *HomePage {
	time.Sleep(time.Second * 2)
	s.Page.ClickByText(btnMasukPassword)

	return &HomePage{Page: s.Page}
}
