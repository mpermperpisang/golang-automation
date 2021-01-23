package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/support"
)

// ValidateUserIsInHomePage : TO DO
func ValidateUserIsInHomePage() error {
	android := androidpages.OnboardingPage{Page: support.DeviceAction}

	android.WaitingMulaiButton().
		ClickMulai().
		ValidateHomePage()

	return nil
}
