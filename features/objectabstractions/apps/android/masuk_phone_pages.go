package androidpages

import (
	"os"
	"time"

	android "github.com/golang-automation/features/helper/apps/android/action"
)

/*InputPhonePage : page object input phone number*/
type InputPhonePage struct {
	Page android.Page
}

var (
	fieldPhone = "//*[@resource-id='app:id/textFieldAV_textField']"
	btnLanjut  = "Lanjut"
)

/*InputPhone : input phone number*/
func (s *InputPhonePage) InputPhone() *InputPhonePage {
	s.Page.SendKeysByXpath(fieldPhone, os.Getenv("USER_PHONE_NUMBER"))

	return &InputPhonePage{Page: s.Page}
}

/*ClickLanjut : click next to input password*/
func (s *InputPhonePage) ClickLanjut() *InputPasswordPage {
	time.Sleep(time.Second * 2)
	s.Page.ClickByText(btnLanjut)

	return &InputPasswordPage{Page: s.Page}
}
