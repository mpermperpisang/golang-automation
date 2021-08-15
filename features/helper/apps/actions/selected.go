package actions

import "github.com/golang-automation/features/helper"

func (s Page) IsElementSelectedByXpath(locator string) bool {
	element, err := s.device().FindByXPath(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByButton(locator string) bool {
	element, err := s.device().FindByButton(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByClass(locator string) bool {
	element, err := s.device().FindByClass(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByID(locator string) bool {
	element, err := s.device().FindByID(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByLabel(locator string) bool {
	element, err := s.device().FindByLabel(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByLink(locator string) bool {
	element, err := s.device().FindByLink(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByName(locator string) bool {
	element, err := s.device().FindByName(locator).Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByText(locator string) bool {
	element, err := s.device().FindByXPath("//*[@text='" + locator + "']").Selected()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByContainsText(locator string) bool {
	element, err := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Selected()
	helper.LogPanicln(err)

	return element
}
