package androidsteps

import (
	"os"
	"time"

	"github.com/golang-automation/features/helper/android"
	"github.com/golang-automation/features/objectabstractions/androidpage"
)

/*GoToLogin is function to go to login page*/
func GoToLogin() error {
	android.ClickByXPath(androidpage.BtnMulai)
	android.ClickByText(androidpage.BtnMasuk)

	return nil
}

/*UserLogin is function to input username and password for android*/
func UserLogin() error {
	android.InputText(androidpage.FieldPhone, os.Getenv("USER_PHONE_NUMBER"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpage.BtnLanjut)
	android.InputText(androidpage.FieldPassword, os.Getenv("USER_PHONE_PASSWORD"))
	time.Sleep(time.Second * 2)
	android.ClickByText(androidpage.BtnMasukPassword)

	return nil
}

/*ValidateAndroidLoggedUser is function to validate user has logged successfully*/
func ValidateAndroidLoggedUser() error {
	onboardingHomePage()
	android.IsElementPresent(androidpage.LabelUsername)

	return nil
}

func onboardingHomePage() error {
	android.ClickByText(androidpage.BtnMengerti)
	android.ClickByText(androidpage.BtnNantiSaja)

	return nil
}
