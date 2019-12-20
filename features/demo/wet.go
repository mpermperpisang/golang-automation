package demo

import (
	"log"
)

func wet() error {
	if "200" != "500" {
		log.Fatalln("Status tidak sama")
	}

	if "Pisang" != "Banana" {
		log.Fatalln("Nama tidak sama")
	}

	if 1 != 2 {
		log.Fatalln("Nilai tidak sama")
	}

	if 1 == 1 {
		log.Fatalln("Nilainya sama")
	}

	if "Pisang" != "Pisang" {
		log.Fatalln("Namanya sama")
	}

	return nil
}
