package helper

import (
	"log"
	"reflect"
)

// AssertEqual : must equal
func AssertEqual(exp interface{}, act interface{}, err interface{}) error {
	if act != exp {
		log.Panicln(err)
	}

	return nil
}

// AssertNotEqual : must not equal
func AssertNotEqual(exp interface{}, act interface{}, err interface{}) error {
	if act == exp {
		log.Panicln(err)
	}

	return nil
}

// AssertString : must string
func AssertString(exp interface{}, act interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "string" {
		log.Panicln(err)
	}

	return nil
}

// AssertInt : must integer
func AssertInt(exp interface{}, act interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "int" {
		log.Panicln(err)
	}

	return nil
}

// AssertFloat64 : must float64
func AssertFloat64(exp interface{}, act interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "float64" {
		log.Panicln(err)
	}

	return nil
}

// AssertBool : must boolean
func AssertBool(exp interface{}, act interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "bool" {
		log.Panicln(err)
	}

	return nil
}
