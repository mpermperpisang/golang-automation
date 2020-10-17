package helper

// Contains : check array contain value
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
