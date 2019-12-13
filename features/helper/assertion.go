package helper

import (
	"log"
)

/*AssertEqual function for assertion*/
func AssertEqual(Expect interface{}, Actual interface{}, ErrMsg interface{}) error {
	if Actual != Expect {
		log.Panicln(ErrMsg)
	}

	return nil
}

/*AssertNotEqual function for assertion*/
func AssertNotEqual(Expect interface{}, Actual interface{}, ErrMsg interface{}) error {
	if Actual == Expect {
		log.Panicln(ErrMsg)
	}

	return nil
}
