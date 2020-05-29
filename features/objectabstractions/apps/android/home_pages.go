package androidpages

import (
	android "github.com/golang-automation/features/helper/apps/android/action"
)

/*HomePage : page object home*/
type HomePage struct {
	Page android.Page
}

var (
	btnMasuk      = "Masuk"
	btnMengerti   = "Mengerti"
	btnNantiSaja  = "Nanti Saja"
	labelUsername = "//*[@resource-id='app:id/tvUsername']"
)

/*VerifyUserIsLogged : validate logged user*/
func (s *HomePage) VerifyUserIsLogged() *HomePage {
	s.Page.IsElementVisibleByXpath(labelUsername)

	return &HomePage{Page: s.Page}
}

/*ClickOnboardingHomePage : click Mengerti and Nanti Saja button*/
func (s *HomePage) ClickOnboardingHomePage() *HomePage {
	s.Page.ClickByText(btnMengerti)
	s.Page.ClickByText(btnNantiSaja)

	return &HomePage{Page: s.Page}
}

/*ClickBtnMasuk : click Masuk button*/
func (s *HomePage) ClickBtnMasuk() *InputPhonePage {
	s.Page.ClickByText(btnMasuk)

	return &InputPhonePage{Page: s.Page}
}
