package helper

/*RegexReadENV : regex to read env variable*/
func RegexReadENV() string {
	return `ENV:([a-zA-Z0-9_]+)`
}

/*RegexInt : regex to read integer*/
func RegexInt() string {
	return `\d+`
}

/*RegexBaseURL : regex to read base url*/
func RegexBaseURL() string {
	return `preprod?`
}

/*RegexTag : regex to read scenario tag from command line*/
func RegexTag() string {
	return `@[a-zA-Z0-9-]*`
}
