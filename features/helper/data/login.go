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

// GetUsername : get data for username
func GetUsername() string {
	return login.username
}

// SetPassword : set data for password
func SetPassword(password string) {
	login.password = password
}

// GetPassword : get data for password
func GetPassword() string {
	return login.password
}
