package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/logrusorgru/aurora"
	"github.com/yalp/jsonpath"
)

var HttpResponse *http.Response
var BaseURL, readENV string
var ResponseBody []byte
var readEndpoint bool
var httpResponse, AccessToken interface{}

func envreader(env string) error {
	readEndpoint = strings.HasPrefix(env, "ENV")
	readENV = strings.TrimPrefix(env, "ENV:")

	return nil
}

/*BaseAPI is function to set base url for API*/
func BaseAPI(base string) error {
	envreader(base)

	if readEndpoint {
		BaseURL = os.Getenv(readENV)
	}

	return nil
}

/*Authentication is function to get access token*/
func Authentication(account string) error {
	envLogin := strings.ToUpper(account)
	username := os.Getenv(envLogin + "_USERNAME")
	password := os.Getenv(envLogin + "_PASSWORD")
	readURL := BaseURL + os.Getenv("AUTH_ENDPOINT")
	var number = regexp.MustCompile(`\d+`).FindString(BaseURL)

	body := []byte(
		`{
			"grant_type": "password",
			"username": "` + username + `",
			"password": "` + password + `",
			"client_id": "` + os.Getenv("API_CLIENT_ID"+"_"+number) + `",
			"client_secret": "` + os.Getenv("API_CLIENT_SECRET"+"_"+number) + `",
			"scope": "public user"
		}`)

	client := &http.Client{}
	httpRequest, _ := http.NewRequest("POST", readURL, bytes.NewBuffer(body))

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	HTTPResponse, _ := client.Do(httpRequest)
	ResponseBody, _ := ioutil.ReadAll(HTTPResponse.Body)
	http, _ := jsonpath.Prepare("$.access_token")

	if err := json.Unmarshal(ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	AccessToken, _ = http(httpResponse)

	return nil
}

/*RetrieveAPI is function to send request*/
func RetrieveAPI(verbose string, endpoint string, body string) error {
	envreader(endpoint)
	readVerbose := strings.ToUpper(verbose)

	if readEndpoint {
		endpoint = os.Getenv(readENV)
	}

	requestBody := []byte(body)
	readURL := BaseURL + endpoint
	client := &http.Client{}
	httpRequest, _ := http.NewRequest(readVerbose, readURL, bytes.NewBuffer(requestBody))

	if AccessToken != nil {
		httpRequest.Header.Add("Authorization", "Bearer "+AccessToken.(string))
	}

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	HttpResponse, _ = client.Do(httpRequest)
	ResponseBody, _ = ioutil.ReadAll(HttpResponse.Body)

	return nil
}

/*ResponseStatusCodeAPI is function to get response code*/
func ResponseStatusCodeAPI(response int) error {
	actualCode := HttpResponse.StatusCode

	if expectCode := (response); actualCode != expectCode {
		log.Fatalln("actual status code :", aurora.Bold(aurora.Red(actualCode)))
	}

	return nil
}
