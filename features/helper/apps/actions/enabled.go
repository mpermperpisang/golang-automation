package actions

import "github.com/golang-automation/features/helper"

func (s Page) IsElementEnabledByXpath(locator string) bool {
	element, err := s.device().FindByXPath(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByButton(locator string) bool {
	element, err := s.device().FindByButton(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByClass(locator string) bool {
	element, err := s.device().FindByClass(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByID(locator string) bool {
	element, err := s.device().FindByID(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByLabel(locator string) bool {
	element, err := s.device().FindByLabel(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByLink(locator string) bool {
	element, err := s.device().FindByLink(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByName(locator string) bool {
	element, err := s.device().FindByName(locator).Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByText(locator string) bool {
	element, err := s.device().FindByXPath("//*[@text='" + locator + "']").Enabled()
	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByContainsText(locator string) bool {
	element, err := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Enabled()
	helper.LogPanicln(err)

	return element
}
