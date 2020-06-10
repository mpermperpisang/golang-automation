package apisteps

import (
	"github.com/DATA-DOG/godog/gherkin"
	api "github.com/golang-automation/features/helper/api"
	step "github.com/golang-automation/features/stepdefinitions/api/steps"
)

/*AuthenticationAPI : login auth*/
func AuthenticationAPI(account string) error {
	api.AuthenticationAPI(account)

	return nil
}

/*RequestAPIWithoutBody : send request API without body*/
func RequestAPIWithoutBody(verbose string, request string) error {
	api.RequestAPI(verbose, request, "")

	return nil
}

/*RequestAPIWithBody : send request API with body*/
func RequestAPIWithBody(verbose string, request string, body *gherkin.DocString) error {
	api.RequestAPI(verbose, request, body.Content)

	return nil
}

/*ResponseDataType : get data type of response*/
func ResponseDataType(path string, expect string) error {
	step.ResponseDataType(path, expect)

	return nil
}

/*ResponseMatchingValue : matching json value*/
func ResponseMatchingValue(path string, expect string) error {
	step.ResponseMatchingValue(path, expect)

	return nil
}

/*ResponseFindPath : find json key*/
func ResponseFindPath(path string) error {
	step.ResponseFindPath(path)

	return nil
}
