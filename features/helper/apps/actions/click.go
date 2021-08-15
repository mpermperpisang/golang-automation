package actions

func (s Page) ClickByXPath(locator string) error {
	return s.device().FindByXPath(locator).Click()
}

func (s Page) ClickByText(locator string) error {
	return s.device().FindByXPath("//*[@text='" + locator + "']").Click()
}

func (s Page) ClickByContainsText(locator string) error {
	return s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Click()
}

func (s Page) ClickByButton(locator string) error {
	return s.device().FindByButton(locator).Click()
}

func (s Page) ClickByClass(locator string) error {
	return s.device().FindByClass(locator).Click()
}

func (s Page) ClickByID(locator string) error {
	return s.device().FindByID(locator).Click()
}

func (s Page) ClickByLabel(locator string) error {
	return s.device().FindByLabel(locator).Click()
}

func (s Page) ClickByLink(locator string) error {
	return s.device().FindByLink(locator).Click()
}

func (s Page) ClickByName(locator string) error {
	return s.device().FindByName(locator).Click()
}
