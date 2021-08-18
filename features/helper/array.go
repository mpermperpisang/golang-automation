package helper

func CheckEmpty(array []int) int {
	if len(array) < 1 {
		return 1
	} else {
		return array[0]
	}
}
