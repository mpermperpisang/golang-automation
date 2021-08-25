package actions

import "golang-automation/features/helper"

func (s Page) GetAttributeByXpath(locator string, attr string) string {
	element, err := s.device().FindByXPath(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByButton(locator string, attr string) string {
	element, err := s.device().FindByButton(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByClass(locator string, attr string) string {
	element, err := s.device().FindByClass(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByID(locator string, attr string) string {
	element, err := s.device().FindByID(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByLabel(locator string, attr string) string {
	element, err := s.device().FindByLabel(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByLink(locator string, attr string) string {
	element, err := s.device().FindByLink(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByName(locator string, attr string) string {
	element, err := s.device().FindByName(locator).Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByText(locator string, attr string) string {
	element, err := s.device().FindByXPath("//*[@text='" + locator + "']").Attribute(attr)
	helper.LogPanicln(err)

	return element
}

func (s Page) GetAttributeByContainsText(locator string, attr string) string {
	element, err := s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Attribute(attr)
	helper.LogPanicln(err)

	return element
}
