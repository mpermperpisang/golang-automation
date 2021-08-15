package stepdefinitions

import (
	formats "github.com/golang-automation/features/helper/formats/web/desktop"
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
	"github.com/golang-automation/features/supports/structs"
)

func VisitLink(url string) error {
	home := structs.WebDriverConnect()

	home.GoToURL(formats.CompleteLink(url))

	return nil
}

func ClickButton(button string) error {
	home := desktoppages.HomePage{Page: structs.WebDriverConnect()}

	home.ClickButton(button)

	return nil
}

func ValidateLoginWarning() error {
	home := desktoppages.HomePage{Page: structs.WebDriverConnect()}

	home.ValidateWarning()

	return nil
}
