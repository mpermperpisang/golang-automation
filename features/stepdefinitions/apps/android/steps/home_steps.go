package step

import (
	android "github.com/golang-automation/features/helper/apps/android/action"
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
)

/*VerifyUserIsLogged : validate logged user*/
func VerifyUserIsLogged() error {
	android.IsElementVisibleByXpath(androidpages.LabelUsername)

	return nil
}

/*ClickOnboardingHomePage : click Mengerti and Nanti Saja button*/
func ClickOnboardingHomePage() error {
	android.ClickByText(androidpages.BtnMengerti)
	android.ClickByText(androidpages.BtnNantiSaja)

	return nil
}

/*ClickBtnMasuk : click Masuk button*/
func ClickBtnMasuk() error {
	android.ClickByText(androidpages.BtnMasuk)

	return nil
}
