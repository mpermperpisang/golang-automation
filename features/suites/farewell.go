package suites

import (
	"github.com/cucumber/godog"
	"github.com/golang-automation/features/stepdefinitions"
)

// FarewellGreeting : suites for API
func FarewellGreeting(s *godog.Suite) {
	s.Step(`^waktu terus berjalan$`, stepdefinitions.FarewellOne)
	s.Step(`^tanpa sadar diriku sampai di ujung waktu$`, stepdefinitions.FarewellTwo)
	s.Step(`^menoleh ke belakang$`, stepdefinitions.FarewellThree)
	s.Step(`^tak ada satupun yang terasa sia - sia$`, stepdefinitions.FarewellFour)
	s.Step(`^bukan air mata pilu yang terbit di ujung mata$`, stepdefinitions.FarewellFive)
	s.Step(`^haru$`, stepdefinitions.FarewellSix)
	s.Step(`^berdiri di ujung waktu$`, stepdefinitions.FarewellSeven)
	s.Step(`^ku ucapkan selamat tinggal$`, stepdefinitions.FarewellEight)
	s.Step(`^terakhir dariku$`, stepdefinitions.FarewellNine)
	s.Step(`^bukan akhir cerita$`, stepdefinitions.FarewellTen)
	s.Step(`^melainkan awal mula bagi cerita setiap kita$`, stepdefinitions.FarewellEleven)
}
