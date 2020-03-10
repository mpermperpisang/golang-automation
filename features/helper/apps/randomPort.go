package appshelper

import (
	"math/rand"
	"time"
)

/*RandomPort fungsi untuk mengacak appium port*/
func RandomPort() int {
	rand.Seed(time.Now().UnixNano())

	return 1000 + rand.Intn(9999-1000)
}
