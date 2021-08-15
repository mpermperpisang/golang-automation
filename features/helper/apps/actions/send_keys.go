package actions

func (s Page) SendKeysByXpath(locator string, text string) error {
	return s.device().FindByXPath(locator).SendKeys(text)
}

func (s Page) SendKeysByButton(locator string, text string) error {
	return s.device().FindByButton(locator).SendKeys(text)
}

func (s Page) SendKeysByClass(locator string, text string) error {
	return s.device().FindByClass(locator).SendKeys(text)
}

func (s Page) SendKeysByID(locator string, text string) error {
	return s.device().FindByID(locator).SendKeys(text)
}

func (s Page) SendKeysByLabel(locator string, text string) error {
	return s.device().FindByLabel(locator).SendKeys(text)
}

func (s Page) SendKeysByLink(locator string, text string) error {
	return s.device().FindByLink(locator).SendKeys(text)
}

func (s Page) SendKeysByName(locator string, text string) error {
	return s.device().FindByName(locator).SendKeys(text)
}

func (s Page) SendKeysByText(locator string, text string) error {
	return s.device().FindByXPath("//*[@text='" + locator + "']").SendKeys(text)
}

func (s Page) SendKeysByContainsText(locator string, text string) error {
	return s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").SendKeys(text)
}
