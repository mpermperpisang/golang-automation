package androidsteps

import (
	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

func verifyUserIsLogged() error {
	android.IsElementVisibleByXpath(androidpages.LabelUsername)

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
