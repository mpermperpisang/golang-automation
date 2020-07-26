package appsaction

// SendKeysByXpath : input text by Xpath selector
func (s *Page) SendKeysByXpath(locator string, text string) error {
	return s.device().FindByXPath(locator).SendKeys(text)
}

// SendKeysByButton : input text by button
func (s *Page) SendKeysByButton(locator string, text string) error {
	return s.device().FindByButton(locator).SendKeys(text)
}

// SendKeysByClass : input text by class
func (s *Page) SendKeysByClass(locator string, text string) error {
	return s.device().FindByClass(locator).SendKeys(text)
}

// SendKeysByID : input text by ID
func (s *Page) SendKeysByID(locator string, text string) error {
	return s.device().FindByID(locator).SendKeys(text)
}

// SendKeysByLabel : input text by label
func (s *Page) SendKeysByLabel(locator string, text string) error {
	return s.device().FindByLabel(locator).SendKeys(text)
}

// SendKeysByLink : input text by link
func (s *Page) SendKeysByLink(locator string, text string) error {
	return s.device().FindByLink(locator).SendKeys(text)
}

// SendKeysByName : input text by class name
func (s *Page) SendKeysByName(locator string, text string) error {
	return s.device().FindByName(locator).SendKeys(text)
}

// SendKeysByText : input text by text in xpath
func (s *Page) SendKeysByText(locator string, text string) error {
	return s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").SendKeys(text)
}
