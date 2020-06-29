package step

import (
	"encoding/json"

	api "github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/assertions"
	"github.com/golang-automation/features/helper/errors"
	"github.com/golang-automation/features/helper/messages"
	"github.com/yalp/jsonpath"
)

// JSONValue : struct for set json value
type JSONValue struct {
	variable interface{}
}

var jsonResponse, actualResult interface{}

func decryptJSONResponse() error {
	err := json.Unmarshal(api.ResponseBody, &jsonResponse)
	errors.LogPanicln(err)

	return nil
}

// ResponseFindPath : find path of response API
func ResponseFindPath(path string) error {
	decryptJSONResponse()

	countpath, _ := jsonpath.Read(jsonResponse, path)
	err := len(countpath.([]interface{}))
	errors.LogPanicln(err)

	return nil
}

func getJSONValue(path string) {
	decryptJSONResponse()

	HTTPJson, _ := jsonpath.Prepare(path)
	actualResult, _ = HTTPJson(jsonResponse)
}

// ResponseMatchingValue : find and matching path value of response API
func ResponseMatchingValue(path string, expectResult string) error {
	getJSONValue(path)

	assertions.AssertEqual(expectResult, actualResult, messages.NotMatchValue(actualResult))

	return nil
}

// ResponseDataType : find and matching path value with data type
func ResponseDataType(path string, expectType string) error {
	var actualType string

	decryptJSONResponse()

	HTTPJson, _ := jsonpath.Prepare(path)
	actualResult, _ := HTTPJson(jsonResponse)

	assertions.AssertEqual(expectType, actualType, messages.NotMatchDataType(actualResult.(string)))

	return nil
}
