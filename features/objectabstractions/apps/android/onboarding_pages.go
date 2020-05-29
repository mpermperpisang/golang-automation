package androidpages

import (
	android "github.com/golang-automation/features/helper/apps/android/action"
)

/*OnboardingPage : page object onboarding*/
type OnboardingPage struct {
	Page android.Page
}

var (
	btnMulai = "//androidx.appcompat.widget.AppCompatTextView[contains(@text, 'Mulai')]"
)

/*ClickBtnMulai : click Mulai button*/
func (s *OnboardingPage) ClickBtnMulai() *HomePage {
	s.Page.ClickByXPath(btnMulai)

	return &HomePage{Page: s.Page}
}
