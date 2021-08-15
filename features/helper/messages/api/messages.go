package messages

import (
	"fmt"
)

func ResponseError(expected, actual interface{}) string {
	return fmt.Sprintf("Expected : %v. Actual : %v", expected, actual)
}
