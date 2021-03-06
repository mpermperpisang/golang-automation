package mobilepages

import (
	"github.com/golang-automation/features/helper/data"
	webaction "github.com/golang-automation/features/helper/web/action"
)

// LoginPage : page object login
type LoginPage struct {
	Page webaction.Page
}

var (
	fieldUsername = "user_session_username"
	fieldPassword = "user_session_password"
	btnLogin      = ".js-btn-menu-login.js-tfa-required-button"
)

// InputUsername : input username
func (s LoginPage) InputUsername() *LoginPage {
	s.Page.SendKeysByID(fieldUsername, data.Username())

	return &LoginPage{Page: s.Page}
}

// InputPassword : input password
func (s LoginPage) InputPassword() *LoginPage {
	s.Page.SendKeysByID(fieldPassword, data.Password())

	return &LoginPage{Page: s.Page}
}

// ClickLogin : click Login button
func (s LoginPage) ClickLogin() *OTPPage {
	s.Page.ClickByCSS(btnLogin)

	return &OTPPage{Page: s.Page}
}
