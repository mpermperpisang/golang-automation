package data

type loginData struct {
	username string
	password string
}

var login loginData

// SetUsername : set data for username
func SetUsername(username string) {
	login.username = username
}

// Username : get data for username
func Username() string {
	return login.username
}

// SetPassword : set data for password
func SetPassword(password string) {
	login.password = password
}

// Password : get data for password
func Password() string {
	return login.password
}
