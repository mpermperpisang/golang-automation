package apisteps

import (
	"github.com/cucumber/gherkin-go"
	api "github.com/golang-automation/features/helper/api"
)

// RequestAPIWithoutBody : send request API without body
func RequestAPIWithoutBody(verbose string, request string) error {
	return api.RequestAPI(verbose, request, "")
}

// RequestAPIWithBody : send request API with body
func RequestAPIWithBody(verbose string, request string, body *gherkin.DocString) error {
	return api.RequestAPI(verbose, request, body.Content)
}
