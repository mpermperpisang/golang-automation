package demo

import (
	"fmt"
	"log"
)

func equalCheck(expect string, actual string) error {
	if expect != actual {
		log.Fatal(expect + " tidak sesuai dengan " + actual)
	} else {
		fmt.Println(expect + " dan " + actual + " ternyata sama")
	}

	return nil
}

func dry() error {
	equalCheck("200", "500")
	equalCheck("Pisang", "Banana")
	equalCheck("Pisang", "Pisang")

	return nil
}
