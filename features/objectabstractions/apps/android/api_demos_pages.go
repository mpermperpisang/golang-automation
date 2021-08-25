package pages

import (
	"golang-automation/features/helper"
	appsaction "golang-automation/features/helper/apps/actions"
	messages "golang-automation/features/helper/messages/apps/android"
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
