package actions

import "golang-automation/features/helper"

func (s Page) IsElementSelectedByXpath(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByButton(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByButton(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByClass(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByClass(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByID(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByID(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByLabel(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByLabel(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByLink(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByLink(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByName(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByName(locator).Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByText(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath("//*[@text='" + locator + "']").Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}

func (s Page) IsElementSelectedByContainsText(locator string, timeout ...int) bool {
	max := helper.CheckEmpty(timeout)

	for loop := 0; loop <= max; loop++ {
		element, err = s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Selected()

		helper.WaitElementWithTimeout(element)
	}

	helper.LogPanicln(err)

	return element
}
