package formats

import "fmt"

func CompleteLink(url string) string {
	return fmt.Sprintf("https://m.%s/", url)
}
