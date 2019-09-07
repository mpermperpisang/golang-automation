package stepdefinitions

import (
	"github.com/DATA-DOG/godog"
	"github.com/golang-automation/features/helper/apps/android"
	"github.com/golang-automation/features/objectabstractions"
)

/*GoToAndroidLogin is function to go to logn page*/
func GoToAndroidLogin() error {
	// there is an error framework, need to find out later
	android.FindElementByXpath(objectabstractions.BtnStart).Click()
	android.FindElementByButton(objectabstractions.BtnLoginAndroid).Click()
	return nil
}

/*InputAndroidLogin is function to input username and password for android*/
func InputAndroidLogin() error {
	return godog.ErrPending
}

/*ValidateAndroidLoggedUser is function to validate user has logged successfully*/
func ValidateAndroidLoggedUser() error {
	return godog.ErrPending
}
