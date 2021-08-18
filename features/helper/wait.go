package helper

import (
	"time"
)

func WaitWithTimeout(timeout time.Duration) {
	time.Sleep(timeout * time.Second)
}

func WaitElementWithTimeout(element bool) {
	if !element {
		WaitWithTimeout(1)
	}
}
