package messages

import "fmt"

func NotExistFileFormat(fileformat string) string {
	return fmt.Sprintf("Not exist file format : %s", fileformat)
}
