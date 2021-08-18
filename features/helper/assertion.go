package helper

import (
	"fmt"
	"reflect"
	"testing"

	apimessages "github.com/golang-automation/features/helper/messages/api"
	"github.com/stretchr/testify/assert"
)

type UseAssertion struct {
	testing testing.T
}

func (t *UseAssertion) assertNew() *assert.Assertions {
	return assert.New(&t.testing)
}

func (t *UseAssertion) AssertEqual(expected, actual, err interface{}) bool {
	assert := t.assertNew().Equal(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)

	if !assert {
		LogPanicln(apimessages.ResponseError(expected, actual))
	}

	return true
}

func (t *UseAssertion) AssertNotEqual(expected, actual, err interface{}) bool {
	assert := t.assertNew().NotEqual(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)

	if !assert {
		LogPanicln(apimessages.ResponseError(expected, actual))
	}

	return true
}

func AssertString(expected, actual, err interface{}) bool {
	if reflect.TypeOf(actual).String() != "string" {
		LogPanicln(err)
	}

	return true
}

func AssertInt(actual, err interface{}) bool {
	if reflect.TypeOf(actual).String() != "int" {
		LogPanicln(err)
	}

	return true
}

func AssertFloat64(expected, actual, err interface{}) bool {
	if reflect.TypeOf(actual).String() != "float64" {
		LogPanicln(err)
	}

	return true
}

func AssertBool(expected, actual, err interface{}) bool {
	if reflect.TypeOf(actual).String() != "bool" {
		LogPanicln(err)
	}

	return true
}

func (t *UseAssertion) AssertEmpty(expected, err interface{}) bool {
	return t.assertNew().Empty(expected, err)
}

func (t *UseAssertion) AssertContains(expected, actual, err interface{}) bool {
	return t.assertNew().Contains(expected, actual, err)
}

func (t *UseAssertion) AssertNotContains(expected, actual, err interface{}) bool {
	return t.assertNew().NotContains(expected, actual, err)
}

func (t *UseAssertion) AssertZero(expected, err interface{}) bool {
	return t.assertNew().Zero(expected, err)
}

func (t *UseAssertion) AssertNotZero(expected, err interface{}) bool {
	return t.assertNew().NotZero(expected, err)
}

func (t *UseAssertion) AssertTrue(expected bool, err interface{}) bool {
	return t.assertNew().True(expected, err)
}

func (t *UseAssertion) AssertFalse(expected bool, err interface{}) bool {
	return t.assertNew().False(expected, err)
}

func (t *UseAssertion) AssertSame(expected, actual, err interface{}) bool {
	return t.assertNew().Same(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)
}

func (t *UseAssertion) AssertNil(expected, err interface{}) bool {
	return t.assertNew().Nil(expected, err)
}

func (t *UseAssertion) AssertNotNil(expected, err interface{}) bool {
	return t.assertNew().NotNil(expected, err)
}

func (t *UseAssertion) AssertError(log error, err interface{}) bool {
	return t.assertNew().Error(log, err)
}

func (t *UseAssertion) AssertFail(expected string) bool {
	return t.assertNew().Fail(expected)
}

func (t *UseAssertion) AssertGreater(expected, actual, err interface{}) bool {
	return t.assertNew().Greater(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)
}

func (t *UseAssertion) AssertLess(expected, actual, err interface{}) bool {
	return t.assertNew().Less(fmt.Sprintf("%v", expected), fmt.Sprintf("%v", actual), err)
}
