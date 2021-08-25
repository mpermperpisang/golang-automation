package stepdefinitions

import (
	"bytes"
	"encoding/json"

	"github.com/golang-automation-v1/features/helper"
	apihelper "github.com/golang-automation-v1/features/helper/api"
	apimessages "github.com/golang-automation-v1/features/helper/messages/api"
	apisupport "github.com/golang-automation-v1/features/supports/base/api"

	"github.com/cucumber/godog"
	"github.com/yalp/jsonpath"
)

var (
	validate = helper.UseAssertion{}

	jsonResponse interface{}
)

func HitEndpointWithBody(method, url string, body *godog.DocString) error {
	apisupport.ClientDo(apisupport.SendRequest(method, url, bytes.NewBuffer([]byte(apihelper.EnvReader(body)))))

	return nil
}

func HitEndpointWithoutBody(method, url string) error {
	apisupport.ClientDo(apisupport.SendRequest(method, url, nil))

	return nil
}

func ValidateStatus(expected int) error {
	actual := apisupport.Response.StatusCode

	validate.AssertEqual(expected, actual, apimessages.ResponseError(expected, actual))

	return nil
}

func ValidateValue(path, expected string) error {
	jsonPath, err := jsonpath.Prepare(path)
	helper.LogPanicln(err)

	json.Unmarshal(apisupport.ResponseBody, &jsonResponse)
	assertEqualValue(jsonPath, expected)

	return nil
}

func assertEqualValue(jsonPath jsonpath.FilterFunc, expected string) {
	actual, err := jsonPath(jsonResponse)
	helper.LogPanicln(err)

	validate.AssertEqual(expected, actual, apimessages.ResponseError(expected, actual))
}
