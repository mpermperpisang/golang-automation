package actions

import "github.com/mpermperpisang/golang-automation-v1/features/helper"

func (s Page) IsElementEnabledByXpath(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByButton(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByButton(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByClass(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByClass(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByID(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByID(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByLabel(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByLabel(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByLink(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByLink(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByName(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByName(locator).Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByText(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath("//*[@text='" + locator + "']").Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementEnabledByContainsText(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Enabled()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}
