package stepdefinitions

import (
	pages "github.com/golang-automation-v1/features/objectabstractions/apps/android"
	"github.com/golang-automation-v1/features/supports/structs"

	"github.com/cucumber/godog"
)

func ClickMenu(num *godog.Table) error {
	page := pages.APIDemos{Page: structs.AppsDeviceConnect()}

	for index := 1; index < len(num.Rows); index++ {
		page.ClickMenu(num.Rows[index].Cells[0].Value)
	}

	return nil
}

func ValidateActionBarTabs() error {
	page := pages.APIDemos{Page: structs.AppsDeviceConnect()}

	page.ValidateButton()

	return nil
}
