package data

// LoginData : data for login
type LoginData struct {
	username string
	password string
}

var login LoginData
var username = login.username
var password = login.password

/*SetUsername : set data for username*/
func SetUsername(username string) {
	login.username = username
}

/*GetUsername : get data for username*/
func GetUsername() string {
	return username
}

/*SetPassword : set data for password*/
func SetPassword(password string) {
	login.password = password
}

/*GetPassword : get data for password*/
func GetPassword() string {
	return password
}
