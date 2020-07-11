package helper

import (
	"fmt"
	"log"
)

// LogPanicln : general error message using log.Panicln
func LogPanicln(err interface{}) error {
	if err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}
