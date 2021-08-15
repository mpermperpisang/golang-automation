package actions

func (s Page) FillElementByXpath(locator string, text string) error {
	return s.device().FindByXPath(locator).Fill(text)
}

func (s Page) FillElementByButton(locator string, text string) error {
	return s.device().FindByButton(locator).Fill(text)
}

func (s Page) FillElementByClass(locator string, text string) error {
	return s.device().FindByClass(locator).Fill(text)
}

func (s Page) FillElementByID(locator string, text string) error {
	return s.device().FindByID(locator).Fill(text)
}

func (s Page) FillElementByLabel(locator string, text string) error {
	return s.device().FindByLabel(locator).Fill(text)
}

func (s Page) FillElementByLink(locator string, text string) error {
	return s.device().FindByLink(locator).Fill(text)
}

func (s Page) FillElementByName(locator string, text string) error {
	return s.device().FindByName(locator).Fill(text)
}

func (s Page) FillElementByText(locator string, text string) error {
	return s.device().FindByXPath("//*[@text='" + locator + "']").Fill(text)
}

func (s Page) FillElementByContainsText(locator string, text string) error {
	return s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
}
