package stepdefinitions

import (
	pages "github.com/golang-automation-v1/features/objectabstractions/apps/ios"
	"github.com/golang-automation-v1/features/supports/structs"
)

func OpenIOS() error {
	page := pages.APIDemos{Page: structs.AppsDeviceConnect()}

	page.ValidateButton()

	return nil
}
