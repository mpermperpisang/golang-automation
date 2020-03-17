package step

import (
	"time"

	web "github.com/golang-automation/features/helper/web/action"
	desktoppages "github.com/golang-automation/features/objectabstractions/web/desktop"
)

/*ValidateLoggedUser is function to check if user has login successfully*/
func ValidateLoggedUser() error {
	time.Sleep(time.Second * 10)
	web.IsElementDisplayedByCSS(desktoppages.IconProfile)

	return nil
}
