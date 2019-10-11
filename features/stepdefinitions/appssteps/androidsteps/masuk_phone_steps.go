package androidsteps

import (
	"os"
	"time"

	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/objectabstractions/apps/androidpages"
)

/*formPhone function to input phone then click next to input password*/
func formPhone() error {
	android.InputText(androidpages.FieldPhone, os.Getenv("USER_PHONE_NUMBER"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpages.BtnLanjut)

	return nil
}
