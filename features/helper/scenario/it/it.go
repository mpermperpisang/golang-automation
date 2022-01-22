package it

import (
	"fmt"
)

func ResponseCodeAssert(code string) string {
	return fmt.Sprintf("response code should %s", code)
}
