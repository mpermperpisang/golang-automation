package messages

import "fmt"

func NotExistButton(button string) string {
	return fmt.Sprintf("Not exist button : %s", button)
}

func NotExistLink(link string) string {
	return fmt.Sprintf("Not exist link : %s", link)
}
