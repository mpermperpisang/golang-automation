package androidpages

import (
	android "github.com/golang-automation/features/helper/apps/action"
	"github.com/golang-automation/features/helper/page"
)

// OnboardingPage : page object onboarding
type OnboardingPage struct {
	Page android.Page
}

var (
	btnMulai = "oldButtonAV"
)

// WaitingMulaiButton : waiting for Mulai button to display
func (s *OnboardingPage) WaitingMulaiButton() *OnboardingPage {
	page.ValidateElementWithTimeout(s.Page.IsElementEnabledByID(btnMulai), 5)

	return &OnboardingPage{Page: s.Page}
}

// ClickMulaiButton : click Mulai button
func (s *OnboardingPage) ClickMulaiButton() *HomePage {
	s.Page.ClickByID(btnMulai)

	return &HomePage{Page: s.Page}
}
