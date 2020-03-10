package androidsteps

/*GoToLogin function to go to login page from onboarding page*/
func GoToLogin() error {
	clickBtnMulai()
	clickBtnMasuk()

	return nil
}

/*UserLogin function to fill in and process login form*/
func UserLogin() error {
	formPhone()
	formPassword()

	return nil
}

/*LoggedUser is function to validate user has logged successfully*/
func LoggedUser() error {
	clickOnboardingHomePage()
	verifyUserIsLogged()

	return nil
}
