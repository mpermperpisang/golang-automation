package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func BaseAPI(base string) error {
	envreader(base)

	if readEndpoint {
		BaseURL = os.Getenv(readENV)
	}

	return nil
}

func Authentication(account string) error {
	envLogin := strings.ToUpper(account)
	username := os.Getenv(envLogin + "_USERNAME")
	password := os.Getenv(envLogin + "_PASSWORD")
	body := []byte(
		`{
			"grant_type": "password", 
			"username": "` + username + `", 
			"password": "` + password + `", 
			"client_id": "` + os.Getenv("API_CLIENT_ID") + `", 
			"client_secret": "` + os.Getenv("API_CLIENT_SECRET") + `", 
			"scope": "public user"
		}`)
	readURL := BaseURL + os.Getenv("AUTH_ENDPOINT")
	client := &http.Client{}
	httpRequest, _ := http.NewRequest("POST", readURL, bytes.NewBuffer(body))

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	HTTPResponse, _ := client.Do(httpRequest)
	ResponseBody, _ := ioutil.ReadAll(HTTPResponse.Body)
	http, _ := jsonpath.Prepare("$..access_token")

	if err := json.Unmarshal(ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	AccessToken, _ = http(httpResponse)
	fmt.Println(string(readURL))

	return nil
}

func RetrieveAPI(verbose string, endpoint string, body string) error {
	readVerbose := strings.ToUpper(verbose)
	requestBody := []byte(
		`` + body + ``)

	if readEndpoint {
		endpoint = os.Getenv(readENV)
	}

	readURL := BaseURL + endpoint
	client := &http.Client{}
	httpRequest, _ := http.NewRequest(readVerbose, readURL, bytes.NewBuffer(requestBody))

	// fmt.Println(AccessToken)
	// fmt.Println("Bearer " + AccessToken.([]string))

	// httpRequest.Header.Set("Authorization", "Bearer "+AccessToken.(string))
	// httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	HttpResponse, _ = client.Do(httpRequest)
	ResponseBody, _ = ioutil.ReadAll(HttpResponse.Body)
	fmt.Println(string(ResponseBody))

	return nil
}

func ResponseStatusCodeAPI(response int) error {
	actualCode := HttpResponse.StatusCode

	if expectCode := (response); actualCode != expectCode {
		log.Fatalln("actual status code :", aurora.Bold(aurora.Red(actualCode)))
	}

	return nil
}
