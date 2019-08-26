package stepdefinitions

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/DATA-DOG/godog/gherkin"
	"github.com/golang-automation/features/helper/api"
	"github.com/logrusorgru/aurora"
	"github.com/yalp/jsonpath"
)

var httpResponse interface{}

/*AuthenticationAPI is function to login auth*/
func AuthenticationAPI(account string) error {
	api.AuthenticationAPI(account)

	return nil
}

/*RequestAPIWithoutBody is function to initiate request API without define body in gherkin*/
func RequestAPIWithoutBody(verbose string, request string) error {
	fmt.Println("a")
	api.RetrieveAPI(verbose, request, "")
	fmt.Println("b")

	return nil
}

/*RequestAPIWithBody is function to initiate request API with body*/
func RequestAPIWithBody(verbose string, request string, body *gherkin.DocString) error {
	fmt.Println("1")
	api.RetrieveAPI(verbose, request, body.Content)
	fmt.Println("2")

	return nil
}

/*ResponseFindKey is function to find key of response API*/
func ResponseFindKey(key string) error {
	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	countKey, _ := jsonpath.Read(httpResponse, key)
	if err := len(countKey.([]interface{})); err == 0 {
		log.Fatalln(aurora.Bold(aurora.Red(key + " is " + strconv.Itoa(err))))
	}

	return nil
}

/*ResponseMatchingValue is function to find and matching key value of response API*/
func ResponseMatchingValue(key string, response string) error {
	http, _ := jsonpath.Prepare(key)

	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	actualResult, _ := http(httpResponse)

	if expectResult := response; actualResult != expectResult {
		log.Fatalln("actual result :", aurora.Bold(aurora.Red(actualResult)))
	}

	return nil
}

/*ResponseDataType is function to find and matching key value with data type*/
func ResponseDataType(key string, expectType string) error {
	var actualType string

	http, _ := jsonpath.Prepare(key)

	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	actualResult, _ := http(httpResponse)

	switch actualResult.(type) {
	case int:
		actualType = "integer"
	case float64:
		actualType = "float64"
	case string:
		actualType = "string"
	default:
		actualType = "unknown"
	}

	if actualType != expectType {
		log.Fatalln(aurora.Bold(aurora.Red("actual data type : " + actualType)))
	}

	return nil
}
