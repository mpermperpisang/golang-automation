package appsaction

// ClickByXPath : click element by Xpath selector
func (s Page) ClickByXPath(locator string) error {
	return s.device().FindByXPath(locator).Click()
}

// ClickByText : click element by text in xpath
func (s Page) ClickByText(locator string) error {
	return s.device().FindByXPath("//*[contains(@text, '" + locator + "')]").Click()
}

// ClickByButton : click element by button
func (s Page) ClickByButton(locator string) error {
	return s.device().FindByButton(locator).Click()
}

// ClickByClass : click element by class
func (s Page) ClickByClass(locator string) error {
	return s.device().FindByClass(locator).Click()
}

// ClickByID : click element by class ID
func (s Page) ClickByID(locator string) error {
	return s.device().FindByID(locator).Click()
}

// ClickByLabel : click element by label
func (s Page) ClickByLabel(locator string) error {
	return s.device().FindByLabel(locator).Click()
}

// ClickByLink : click element by link
func (s Page) ClickByLink(locator string) error {
	return s.device().FindByLink(locator).Click()
}

// ClickByName : click element by class name
func (s Page) ClickByName(locator string) error {
	return s.device().FindByName(locator).Click()
}
