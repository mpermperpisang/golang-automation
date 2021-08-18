package stepdefinitions

import (
	komit "github.com/golang-automation/features/objectabstractions/apps/android/komit"
	"github.com/golang-automation/features/supports/structs"
)

var page = komit.OnboardingPage{Page: structs.AppsDeviceConnect()}

func GetOnboarding() error {
	page.ValidateOnboardingIsDisplayed(1)

	return nil
}

func ClickLanjutButton() error {
	page.ClickLanjutButton().
		ValidateOnboardingIsDisplayed(2).
		ClickLanjutButton().
		ValidateOnboardingIsDisplayed(3).
		ClickLanjutButton().
		ValidateOnboardingIsDisplayed(4)

	return nil
}

func ValidateBerlanggananPage() error {
	page.ClickOkeMulaiButton().
		ValidateBerlanggananPage()

	return nil
}
