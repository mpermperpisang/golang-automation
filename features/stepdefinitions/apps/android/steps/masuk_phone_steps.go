package step

import (
	"os"
	"time"

	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*InputPhone : input phone then click next to input password*/
func InputPhone() error {
	android.SendKeysByXpath(androidpages.FieldPhone, os.Getenv("USER_PHONE_NUMBER"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpages.BtnLanjut)

	return nil
}
