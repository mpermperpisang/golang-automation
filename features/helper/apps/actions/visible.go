package actions

import "github.com/golang-automation/features/helper"

func (s Page) IsElementVisibleByXpath(locator string) bool {
	element, err := s.device().FindByXPath(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByButton(locator string) bool {
	element, err := s.device().FindByButton(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByClass(locator string) bool {
	element, err := s.device().FindByClass(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByID(locator string) bool {
	element, err := s.device().FindByID(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByLabel(locator string) bool {
	element, err := s.device().FindByLabel(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByLink(locator string) bool {
	element, err := s.device().FindByLink(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByName(locator string) bool {
	element, err := s.device().FindByName(locator).Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByText(locator string) bool {
	element, err := s.device().FindByXPath("//*[@text='" + locator + "']").Visible()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByContainsText(locator string) bool {
	element, err := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Visible()
	helper.LogPanicln(err)

	return element
}
