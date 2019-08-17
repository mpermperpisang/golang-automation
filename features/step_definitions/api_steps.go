package step_definitions

import (
	"encoding/json"
	"log"

	"github.com/golang-automation/features/helper/api"
	. "github.com/logrusorgru/aurora"
	"github.com/yalp/jsonpath"
)

var httpResponse interface{}

func ResponseMatchingKey(data string) error {
	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatal(Bold(Red(err)))
	}

	if _, err := jsonpath.Read(httpResponse, data); err != nil {
		log.Fatal(Bold(Red(err)))
	}

	return nil
}

func ResponseMatchingValue(data string, response string) error {
	http, _ := jsonpath.Prepare(data)

	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatal(Bold(Red(err)))
	}

	actualResult, _ := http(httpResponse)
	expectResult := response

	if actualResult != expectResult {
		log.Fatalln("actual status code :", Bold(Red(actualResult)))
	}

	return nil
}
