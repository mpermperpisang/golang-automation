package androidsteps

import (
	"os"
	"time"

	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*formPassword function to input password and click enter to login*/
func formPassword() error {
	android.SendKeysByXpath(androidpages.FieldPassword, os.Getenv("USER_PHONE_PASSWORD"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpages.BtnMasukPassword)

	return nil
}
