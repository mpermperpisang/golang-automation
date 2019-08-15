package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/logrusorgru/aurora"
)

var httpResponse *http.Response

func RetrieveAPI(verbose string, request string) error {
	readEndpoint := strings.HasPrefix(request, "ENV")
	readENV := strings.TrimPrefix(request, "ENV:")
	readVerbose := strings.ToUpper(verbose)

	if readEndpoint {
		request = os.Getenv(readENV)
	}

	readURL := os.Getenv("API_BASE_URL") + request
	client := &http.Client{}
	httpRequest, _ := http.NewRequest(readVerbose, readURL, nil)

	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	httpResponse, _ = client.Do(httpRequest)

	return nil
}

func ResponseAPI(response int) error {
	actualCode := httpResponse.StatusCode

	if expectCode := (response); actualCode != expectCode {
		log.Fatalln("actual status code :", Bold(Red(actualCode)))
	} else {
		fmt.Println(Bold(Magenta("SCENARIO IS SUCCESS")))
	}

	return nil
}
