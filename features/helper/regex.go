package helper

func RegexENV() string {
	return `ENV:([a-zA-Z0-9_]+)`
}

func RegexInt() string {
	return `\d+`
}

func RegexTag() string {
	return `@[a-zA-Z0-9-]*`
}
