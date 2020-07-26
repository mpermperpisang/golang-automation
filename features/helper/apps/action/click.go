package appsaction

// ClickByXPath : click element by Xpath selector
func (s *Page) ClickByXPath(locator string) error {
	return s.device().AllByXPath(locator).Click()
}

// ClickByText : click element by text in xpath
func (s *Page) ClickByText(locator string) error {
	return s.device().AllByXPath("//*[contains(@text, '" + locator + "')]").Click()
}

// ClickByButton : click element by button
func (s *Page) ClickByButton(locator string) error {
	return s.device().AllByButton(locator).Click()
}

// ClickByClass : click element by class
func (s *Page) ClickByClass(locator string) error {
	return s.device().AllByClass(locator).Click()
}

// ClickByID : click element by class ID
func (s *Page) ClickByID(locator string) error {
	return s.device().AllByID(locator).Click()
}

// ClickByLabel : click element by label
func (s *Page) ClickByLabel(locator string) error {
	return s.device().AllByLabel(locator).Click()
}

// ClickByLink : click element by link
func (s *Page) ClickByLink(locator string) error {
	return s.device().AllByLink(locator).Click()
}

// ClickByName : click element by class name
func (s *Page) ClickByName(locator string) error {
	return s.device().AllByName(locator).Click()
}
