package pages

import (
	"github.com/golang-automation-v1/features/helper"
	appsaction "github.com/golang-automation-v1/features/helper/apps/actions"
	messages "github.com/golang-automation-v1/features/helper/messages/apps/android"
)

type APIDemos struct {
	Page appsaction.Page
}

var (
	appMenu             = "App"
	actionBarMenu       = "Action Bar"
	actionBarTabsMenu   = "Action Bar Tabs"
	addNewTabButton     = "ADD NEW TAB"
	removeLastButton    = "REMOVE LAST TAB"
	toggleTabModeButton = "TOGGLE TAB MODE"
	removeAllTabsButton = "REMOVE ALL TABS"
)

func (s APIDemos) ClickMenu(menu string) *APIDemos {
	var element string

	switch menu {
	case "App":
		element = appMenu
	case "Action Bar":
		element = actionBarMenu
	case "Action Bar Tabs":
		element = actionBarTabsMenu
	default:
		helper.LogPanicln(messages.NotExistMenu(menu))
	}

	s.Page.ClickByText(element)

	return &APIDemos{Page: s.Page}
}

func (s APIDemos) ValidateButton() *APIDemos {
	s.Page.IsElementVisibleByText(addNewTabButton, 3)
	s.Page.IsElementVisibleByText(removeLastButton)
	s.Page.IsElementVisibleByText(toggleTabModeButton)
	s.Page.IsElementVisibleByText(removeAllTabsButton)

	return &APIDemos{Page: s.Page}
}
