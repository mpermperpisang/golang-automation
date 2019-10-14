package demo

import "fmt"

/*Hello function to try unit test*/
func Hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}

	return fmt.Sprintf("Hello %v!", user)
}
