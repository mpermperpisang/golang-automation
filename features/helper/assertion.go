package helper

import (
	"log"
)

/*AssertEqual function for assertion*/
func AssertEqual(exp interface{}, act interface{}, err interface{}) error {
	if act != exp {
		log.Panicln(err)
	}

	return nil
}

/*AssertNotEqual function for assertion*/
func AssertNotEqual(exp interface{}, act interface{}, err interface{}) error {
	if act == exp {
		log.Panicln(err)
	}

	return nil
}
