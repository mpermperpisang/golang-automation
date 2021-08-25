package actions

import (
	"github.com/golang-automation-v1/features/helper"
)

func (s Page) IsElementVisibleByXpath(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByButton(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByButton(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByClass(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByClass(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByID(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByID(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByLabel(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByLabel(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByLink(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByLink(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByName(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByName(locator).Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByText(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath("//*[@text='" + locator + "']").Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementVisibleByContainsText(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Visible()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}
