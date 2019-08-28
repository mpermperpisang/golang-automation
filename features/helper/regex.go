package helper

/*RegexReadENV is function regex to read env variable*/
func RegexReadENV() string {
	return `ENV:([a-zA-Z0-9_]+)`
}
