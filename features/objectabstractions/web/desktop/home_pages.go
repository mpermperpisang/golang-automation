package desktoppages

import (
	"time"

	webaction "github.com/golang-automation/features/helper/web/action"
)

/*HomePage : page object home*/
type HomePage struct {
	Page webaction.Page
}

var (
	iconProfile = ".c-dropdown--profile"
)

/*ValidateLoggedUser : validate user has login successfully*/
func (s *HomePage) ValidateLoggedUser() error {
	time.Sleep(time.Second * 10)
	s.Page.IsElementDisplayedByCSS(iconProfile)

	return nil
}
