package appsaction

// FillElementByXpath : fill element by Xpath selector
func (s Page) FillElementByXpath(locator string, text string) error {
	return s.device().FindByXPath(locator).Fill(text)
}

// FillElementByButton : fill element by button
func (s Page) FillElementByButton(locator string, text string) error {
	return s.device().FindByButton(locator).Fill(text)
}

// FillElementByClass : fill element by class
func (s Page) FillElementByClass(locator string, text string) error {
	return s.device().FindByClass(locator).Fill(text)
}

// FillElementByID : fill element by ID
func (s Page) FillElementByID(locator string, text string) error {
	return s.device().FindByID(locator).Fill(text)
}

// FillElementByLabel : fill element by label
func (s Page) FillElementByLabel(locator string, text string) error {
	return s.device().FindByLabel(locator).Fill(text)
}

// FillElementByLink : fill element by link
func (s Page) FillElementByLink(locator string, text string) error {
	return s.device().FindByLink(locator).Fill(text)
}

// FillElementByName : fill element by class name
func (s Page) FillElementByName(locator string, text string) error {
	return s.device().FindByName(locator).Fill(text)
}

// FillElementByText : fill element by Xpath selector
func (s Page) FillElementByText(locator string, text string) error {
	return s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Fill(text)
}
