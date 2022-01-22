package formats

import "fmt"

func Screenshots() string {
	return "%s - %s.png"
}

func Logs() string {
	return "%s - %s.log"
}

func TestPath(path, platform string) string {
	return fmt.Sprintf("/test/%s/%s/", path, platform)
}
