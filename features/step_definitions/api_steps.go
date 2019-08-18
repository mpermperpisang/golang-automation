package step_definitions

import (
	"encoding/json"
	"log"

	"github.com/golang-automation/features/helper/api"
	. "github.com/logrusorgru/aurora"
	"github.com/yalp/jsonpath"
)

var httpResponse interface{}

func ResponseFindKey(key string) error {
	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(Bold(Red(err)))
	}

	if _, err := jsonpath.Read(httpResponse, key); err != nil {
		log.Fatalln(Bold(Red(err)))
	}

	return nil
}

func ResponseMatchingValue(key string, response string) error {
	http, _ := jsonpath.Prepare(key)

	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(Bold(Red(err)))
	}

	actualResult, _ := http(httpResponse)
	expectResult := response

	if actualResult != expectResult {
		log.Fatalln("actual status code :", Bold(Red(actualResult)))
	}

	return nil
}
