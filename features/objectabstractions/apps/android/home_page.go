package androidpages

import (
	android "github.com/golang-automation/features/helper/apps/action"
	"github.com/golang-automation/features/helper/page"
)

// HomePage : page object home
type HomePage struct {
	Page android.Page
}

var (
	btnMasuk      = "Masuk"
	btnMengerti   = "Mengerti"
	btnNantiSaja  = "Nanti Saja"
	labelUsername = "//*[@resource-id='app:id/tvUsername']"
)

// ValidateLoggedUser : validate logged user
func (s *HomePage) ValidateLoggedUser() *HomePage {
	s.Page.IsElementVisibleByXpath(labelUsername)

	return &HomePage{Page: s.Page}
}

// ClickMengerti : click Mengerti button
func (s *HomePage) ClickMengerti() *HomePage {
	s.Page.ClickByText(btnMengerti)

	return &HomePage{Page: s.Page}
}

// ClickNantiSaja : click Nanti Saja button
func (s *HomePage) ClickNantiSaja() *HomePage {
	s.Page.ClickByText(btnNantiSaja)

	return &HomePage{Page: s.Page}
}

// ClickMasuk : click Masuk button
func (s *HomePage) ClickMasuk() *InputPhonePage {
	s.Page.ClickByText(btnMasuk)

	return &InputPhonePage{Page: s.Page}
}

// ValidateHomePage : validate user is in home page
func (s *HomePage) ValidateHomePage() *HomePage {
	page.ValidateElementWithTimeout(s.Page.IsElementVisibleByText(btnMasuk), 5)

	return &HomePage{Page: s.Page}
}
