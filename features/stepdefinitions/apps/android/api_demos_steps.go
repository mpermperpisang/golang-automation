package stepdefinitions

import (
	"github.com/cucumber/godog"
	pages "github.com/golang-automation/features/objectabstractions/apps/android"
	"github.com/golang-automation/features/supports/structs"
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
