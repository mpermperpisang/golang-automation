package apisteps

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/golang-automation/features/helper/api"
	"github.com/logrusorgru/aurora"
	"github.com/yalp/jsonpath"
)

var jsonResponse interface{}

/*AuthenticationAPI is function to login auth*/
func AuthenticationAPI(account string) error {
	api.AuthenticationAPI(account)

	return nil
}

/*RequestAPIWithoutBody is function to initiate request API without define body in gherkin*/
func RequestAPIWithoutBody(verbose string, request string) error {
	api.RetrieveAPI(verbose, request, "")

	return nil
}

/*RequestAPIWithBody is function to initiate request API with body*/
func RequestAPIWithBody(verbose string, request string, body *gherkin.DocString) error {
	api.RetrieveAPI(verbose, request, body.Content)

	return nil
}

/*ResponseFindKey is function to find key of response API*/
func ResponseFindKey(key string) error {
	if err := json.Unmarshal(api.ResponseBody, &jsonResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	countKey, _ := jsonpath.Read(jsonResponse, key)

	if err := len(countKey.([]interface{})); err == 0 {
		log.Fatalln(aurora.Bold(aurora.Red(key + " is " + strconv.Itoa(err))))
	}

	return nil
}

/*ResponseMatchingValue is function to find and matching key value of response API*/
func ResponseMatchingValue(key string, expectResult string) error {
	HTTPJson, _ := jsonpath.Prepare(key)

	if err := json.Unmarshal(api.ResponseBody, &jsonResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	actualResult, _ := HTTPJson(jsonResponse)

	if actualResult != expectResult {
		log.Fatalln("actual result :", aurora.Bold(aurora.Red(actualResult)))
	}

	return nil
}

/*ResponseDataType is function to find and matching key value with data type*/
func ResponseDataType(key string, expectType string) error {
	var actualType string

	HTTPJson, _ := jsonpath.Prepare(key)

	if err := json.Unmarshal(api.ResponseBody, &jsonResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	actualResult, _ := HTTPJson(jsonResponse)

	switch actualResult.(type) {
	case int:
		actualType = "integer"
	case float64:
		actualType = "float64"
	case string:
		actualType = "string"
	default:
		actualType = actualResult.(string)
	}

	if actualType != expectType {
		log.Fatalln(aurora.Bold(aurora.Red("actual data type : " + actualType)))
	}

	return nil
}
