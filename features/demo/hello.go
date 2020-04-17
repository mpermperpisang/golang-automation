package demo

import "fmt"

/*Hello : say hello to user*/
func Hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}

	return fmt.Sprintf("Hello %v!", user)
}
