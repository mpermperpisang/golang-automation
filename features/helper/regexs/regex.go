package regexs

// RegexReadENV : read env variable
func RegexReadENV() string {
	return `ENV:([a-zA-Z0-9_]+)`
}

// RegexInt : read integer
func RegexInt() string {
	return `\d+`
}

// RegexBaseURL : read base url
func RegexBaseURL() string {
	return `preprod?`
}

// RegexTag : read scenario tag from command line
func RegexTag() string {
	return `@[a-zA-Z0-9-]*`
}
