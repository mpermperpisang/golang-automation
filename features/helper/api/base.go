package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var HttpResponse *http.Response
var BaseURL string
var ResponseBody []byte

func BaseAPI(base string) error {
	readEndpoint := strings.HasPrefix(base, "ENV")
	readENV := strings.TrimPrefix(base, "ENV:")

	if readEndpoint {
		BaseURL = os.Getenv(readENV)
	}

	return nil
}

func RetrieveAPI(verbose string, endpoint string) error {
	readEndpoint := strings.HasPrefix(endpoint, "ENV")
	readENV := strings.TrimPrefix(endpoint, "ENV:")
	readVerbose := strings.ToUpper(verbose)

	if readEndpoint {
		endpoint = os.Getenv(readENV)
	}

	readURL := BaseURL + endpoint
	client := &http.Client{}
	httpRequest, _ := http.NewRequest(readVerbose, readURL, nil)

	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	HttpResponse, _ = client.Do(httpRequest)
	ResponseBody, _ = ioutil.ReadAll(HttpResponse.Body)

	return nil
}

func ResponseStatusCodeAPI(response int) error {
	actualCode := HttpResponse.StatusCode

	if expectCode := (response); actualCode != expectCode {
		log.Fatalln("actual status code :", Bold(Red(actualCode)))
	}

	return nil
}
