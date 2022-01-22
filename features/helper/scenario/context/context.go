package context

import (
	"fmt"
)

func EndpointMethod(method string) string {
	return fmt.Sprintf("%s endpoint", method)
}
