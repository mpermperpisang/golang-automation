package step

import (
	"os"
	"time"

	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*InputPassword : input password and click enter to login*/
func InputPassword() error {
	android.SendKeysByXpath(androidpages.FieldPassword, os.Getenv("USER_PHONE_PASSWORD"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpages.BtnMasukPassword)

	return nil
}
