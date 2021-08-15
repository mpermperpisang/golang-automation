package stepdefinitions

import (
	komid "github.com/golang-automation/features/objectabstractions/apps/android/komid"
	"github.com/golang-automation/features/supports/structs"
)

func GetOnboarding() error {
	onboarding := komid.OnboardingPage{Page: structs.AppsDeviceConnect()}

	onboarding.ValidateOnboardingIsDisplayed(1)

	return nil
}

func ClickLanjutButton() error {
	onboarding := komid.OnboardingPage{Page: structs.AppsDeviceConnect()}

	onboarding.ClickLanjutButton().
		ValidateOnboardingIsDisplayed(2).
		ClickLanjutButton().
		ValidateOnboardingIsDisplayed(3).
		ClickLanjutButton().
		ValidateOnboardingIsDisplayed(4)

	return nil
}

func ValidateBerlanggananPage() error {
	subscribe := komid.OnboardingPage{Page: structs.AppsDeviceConnect()}

	subscribe.ClickOkeMulaiButton().
		ValidateBerlanggananPage()

	return nil
}
