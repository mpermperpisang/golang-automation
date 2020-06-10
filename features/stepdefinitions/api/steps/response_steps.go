package step

import (
	"encoding/json"

	"github.com/golang-automation/features/helper"
	api "github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/message"
	"github.com/yalp/jsonpath"
)

/*JSONValue : struct for set json value*/
type JSONValue struct {
	variable interface{}
}

var jsonResponse, actualResult interface{}

func decryptJSONResponse() error {
	err := json.Unmarshal(api.ResponseBody, &jsonResponse)
	helper.LogPanicln(err)

	return nil
}

/*ResponseFindPath : find path of response API*/
func ResponseFindPath(path string) error {
	decryptJSONResponse()

	countpath, _ := jsonpath.Read(jsonResponse, path)
	err := len(countpath.([]interface{}))
	helper.LogPanicln(err)

	return nil
}

func getJSONValue(path string) {
	decryptJSONResponse()

	HTTPJson, _ := jsonpath.Prepare(path)
	actualResult, _ = HTTPJson(jsonResponse)
}

/*ResponseMatchingValue : find and matching path value of response API*/
func ResponseMatchingValue(path string, expectResult string) error {
	getJSONValue(path)

	helper.AssertEqual(expectResult, actualResult, message.NotMatchValue(actualResult))

	return nil
}

/*ResponseDataType : find and matching path value with data type*/
func ResponseDataType(path string, expectType string) error {
	var actualType string

	decryptJSONResponse()

	HTTPJson, _ := jsonpath.Prepare(path)
	actualResult, _ := HTTPJson(jsonResponse)

	helper.AssertEqual(expectType, actualType, message.NotMatchDataType(actualResult.(string)))

	return nil
}
