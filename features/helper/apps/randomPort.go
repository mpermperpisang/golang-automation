package appshelper

import (
	"math/rand"
	"time"
)

/*RandomPort : mengacak appium port*/
func RandomPort() int {
	rand.Seed(time.Now().UnixNano())

	return 1000 + rand.Intn(9999-1000)
}
