package formats

import "fmt"

func Screenshots() string {
	return "%s - %s.png"
}

func Logs() string {
	return "%s - %s.log"
}

func SSPath(platform string) string {
	return fmt.Sprintf("/test/screenshots/%s/", platform)
}

func LogPath(platform string) string {
	return fmt.Sprintf("/test/logs/%s/", platform)
}
