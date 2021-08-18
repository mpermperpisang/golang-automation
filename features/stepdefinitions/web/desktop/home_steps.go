package stepdefinitions

import (
	formats "github.com/golang-automation/features/helper/formats/web/desktop"
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	"github.com/golang-automation/features/supports/structs"
)

func VisitLink(url string) error {
	page := structs.WebDriverConnect()

	page.GoToURL(formats.CompleteLink(url))

	return nil
}

func ClickButton(button string) error {
	page := desktoppages.HomePage{Page: structs.WebDriverConnect()}

	page.ClickButton(button)

	return nil
}

func ValidateLoginWarning() error {
	page := desktoppages.HomePage{Page: structs.WebDriverConnect()}

	page.ValidateWarning()

	return nil
}
