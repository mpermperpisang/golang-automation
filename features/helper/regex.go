package helper

/*RegexReadENV is function regex to read env variable*/
func RegexReadENV() string {
	return `ENV:([a-zA-Z0-9_]+)`
}

/*RegexInt function regex to read integer*/
func RegexInt() string {
	return `\d+`
}

/*RegexBaseURL function regex to read base url*/
func RegexBaseURL() string {
	return `preprod?`
}

/*RegexTag function regex to read scenario tag from command line*/
func RegexTag() string {
	return `@[a-zA-Z0-9-]*`
}
