package androidsteps

import (
	"os"
	"time"

	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/objectabstractions/apps/androidpages"
)

/*formPassword function to input password and click enter to login*/
func formPassword() error {
	android.InputText(androidpages.FieldPassword, os.Getenv("USER_PHONE_PASSWORD"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpages.BtnMasukPassword)

	return nil
}
