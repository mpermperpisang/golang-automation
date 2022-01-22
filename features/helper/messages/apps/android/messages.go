package messages

import "fmt"

func NotExistPage(page int) string {
	return fmt.Sprintf("Not exist page : %d", page)
}

func NotExistMenu(menu string) string {
	return fmt.Sprintf("Not exist menu : %s", menu)
}
