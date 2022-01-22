package pages

import (
	appsaction "github.com/golang-automation-v1/features/helper/apps/actions"
)

type APIDemos struct {
	Page appsaction.Page
}

var (
	computeMenu  = "Compute Sum"
	alertMenu    = "//*[@value='show alert']"
	gestureMenu  = "//*[@value='Test Gesture']"
	crashMenu    = "//*[@value='Crash']"
	calendarMenu = "//*[@value='Check calendar authorized']"
)

func (s APIDemos) ValidateButton() *APIDemos {
	s.Page.IsElementVisibleByName(computeMenu, 3)
	s.Page.IsElementVisibleByXpath(alertMenu)
	s.Page.IsElementVisibleByXpath(gestureMenu)
	s.Page.IsElementVisibleByXpath(crashMenu)
	s.Page.IsElementVisibleByXpath(calendarMenu)

	return &APIDemos{Page: s.Page}
}
