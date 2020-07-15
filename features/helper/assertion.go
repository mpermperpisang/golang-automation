package helper

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertEqual : must equal
func AssertEqual(expect interface{}, actual interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Equal(expect, actual, err)
}

// AssertNotEqual : must not equal
func AssertNotEqual(expect interface{}, actual interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.NotEqual(expect, actual, err)
}

// AssertString : must string
func AssertString(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "string" {
		log.Panicln(err)
	}

	return nil
}

// AssertInt : must integer
func AssertInt(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "int" {
		log.Panicln(err)
	}

	return nil
}

// AssertFloat64 : must float64
func AssertFloat64(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "float64" {
		log.Panicln(err)
	}

	return nil
}

// AssertBool : must boolean
func AssertBool(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(act).String() != "bool" {
		log.Panicln(err)
	}

	return nil
}

// AssertEmpty : must empty
func AssertEmpty(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Empty(expect, err)
}

// AssertContains : must contains something
func AssertContains(expect interface{}, contain interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Contains(expect, contain, err)
}

// AssertNotContains : must not contains something
func AssertNotContains(expect interface{}, contain interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.NotContains(expect, contain, err)
}

// AssertZero : must zero
func AssertZero(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Zero(expect, err)
}

// AssertNotZero : must not zero
func AssertNotZero(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.NotZero(expect, err)
}

// AssertTrue : must true
func AssertTrue(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.True(expect, err)
}

// AssertFalse : must false
func AssertFalse(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.False(expect, err)
}

// AssertSame : must same
func AssertSame(expect interface{}, actual interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Same(expect, actual, err)
}

// AssertNil : must nil
func AssertNil(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Nil(expect, err)
}

// AssertNotNil : must not nil
func AssertNotNil(expect interface{}, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.NotNil(expect, err)
}

// AssertError : must error
func AssertError(log error, err interface{}) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Error(log, err)
}

// AssertFail : must fail
func AssertFail(expect string) *assert.Assertions {
	assert := assert.New(*testing.T)

	return assert.Fail(expect)
}
