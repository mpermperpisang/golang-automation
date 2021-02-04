package desktoppages

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
	btnLogin      = "Login"
	headerPage    = "Silakan masuk ke dalam akun kamu"
	fbLogin       = "Login dengan Facebook"
	googleLogin   = "Login dengan Google"
)

// ValidateLoginPage : page validation
func (s LoginPage) ValidateLoginPage() *HomePage {
	s.Page.IsElementDisplayedByText(headerPage)
	s.Page.IsElementDisplayedByText(fbLogin)
	s.Page.IsElementDisplayedByText(googleLogin)

	return &HomePage{Page: s.Page}
}

// InputUsername : fill in username
func (s LoginPage) InputUsername() *LoginPage {
	s.Page.SendKeysByID(fieldUsername, data.Username())

	return &LoginPage{Page: s.Page}
}

// InputPassword : fill in password
func (s LoginPage) InputPassword() *LoginPage {
	s.Page.SendKeysByID(fieldPassword, data.Password())

	return &LoginPage{Page: s.Page}
}

// ClickLogin : click Login button
func (s LoginPage) ClickLogin() *OTPPage {
	s.Page.ClickByText(btnLogin)

	return &OTPPage{Page: s.Page}
}
