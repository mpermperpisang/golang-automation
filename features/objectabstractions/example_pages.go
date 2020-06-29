package objectabstractions

import (
	webaction "github.com/golang-automation/features/helper/web/action"
)

// ExamplePage : page object example
type ExamplePage struct {
	Page webaction.Page
}

var (
	btnLogin        = "#login_link"
	fieldName       = "/// [@name='name']"
	radioAge        = "/// [@name='age']"
	optionAge       = "/// [@value='25']"
	fieldOccupation = "/// [@name='occupation']"
	fieldLocation   = "/// [@name='suggested_location']"
	btnSubmit       = ".c-btn--green.u-pad-h--2"
)
