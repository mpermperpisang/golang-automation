package androidsteps

import (
	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/objectabstractions/apps/androidpages"
)

func verifyUserIsLogged() error {
	android.IsElementPresent(androidpages.LabelUsername)

	return nil
}

func clickOnboardingHomePage() error {
	android.ClickByText(androidpages.BtnMengerti)
	android.ClickByText(androidpages.BtnNantiSaja)

	return nil
}

/*clickBtnMasuk function to click Masuk button*/
func clickBtnMasuk() error {
	android.ClickByText(androidpages.BtnMasuk)

	return nil
}
