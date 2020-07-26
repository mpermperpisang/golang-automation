package apisteps

import (
	"github.com/cucumber/gherkin-go"
	api "github.com/golang-automation/features/helper/api"
	step "github.com/golang-automation/features/stepdefinitions/api/steps"
)

// AuthenticationAPI : login auth
func AuthenticationAPI(account string) error {
	return api.AuthenticationAPI(account)
}

// RequestAPIWithoutBody : send request API without body
func RequestAPIWithoutBody(verbose string, request string) error {
	return api.RequestAPI(verbose, request, "")
}

// RequestAPIWithBody : send request API with body
func RequestAPIWithBody(verbose string, request string, body *gherkin.DocString) error {
	return api.RequestAPI(verbose, request, body.Content)
}

// ResponseDataType : get data type of response
func ResponseDataType(path string, expect string) error {
	return step.ResponseDataType(path, expect)
}

// ResponseMatchingValue : matching json value
func ResponseMatchingValue(path string, expect string) error {
	return step.ResponseMatchingValue(path, expect)
}

// ResponseFindPath : find json key
func ResponseFindPath(path string) error {
	return step.ResponseFindPath(path)
}
