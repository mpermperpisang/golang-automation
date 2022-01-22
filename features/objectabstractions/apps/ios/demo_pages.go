package pages

import (
	appsaction "github.com/golang-automation-v1/features/helper/apps/actions"
)

type APIDemos struct {
	Page appsaction.Page
}

var (
	computeMenu  = "(//*[@label='Compute Sum'])[2]"
	alertMenu    = "(//*[@label='show alert'])[2]"
	gestureMenu  = "(//*[@label='Test Gesture'])[2]"
	crashMenu    = "(//*[@label='Crash'])[2]"
	calendarMenu = "(//*[@label='Check calendar authorized'])[2]"
)

func (s APIDemos) ValidateButton() *APIDemos {
	s.Page.IsElementVisibleByXpath(computeMenu, 3)
	s.Page.IsElementVisibleByXpath(alertMenu)
	s.Page.IsElementVisibleByXpath(gestureMenu)
	s.Page.IsElementVisibleByXpath(crashMenu)
	s.Page.IsElementVisibleByXpath(calendarMenu)

	return &APIDemos{Page: s.Page}
}
