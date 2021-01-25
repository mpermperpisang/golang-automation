package androidsteps

import (
	androidpages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/support"
)

// UserIsInHomePage : TO DO
func UserIsInHomePage() error {
	android := androidpages.OnboardingPage{Page: support.DeviceAction}

	android.WaitingMulaiButton().
		ClickMulaiButton().
		ValidateHomePage()

	return nil
}
