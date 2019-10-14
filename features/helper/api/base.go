package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang-automation/features/helper"
	"github.com/yalp/jsonpath"
)

/*HTTPResponse variable*/
var HTTPResponse *http.Response

/*BaseURL variable*/
var BaseURL, readENV string

/*ResponseBody variable*/
var ResponseBody []byte
var endpointENV bool

/*AccessToken variable*/
var AccessToken interface{}
var err error

func envreader(env string) error {
	endpointENV = strings.HasPrefix(env, "ENV:")
	readENV = strings.TrimPrefix(env, "ENV:")

	return nil
}

/*BaseAPI is function to set base url for API*/
func BaseAPI(base string) error {
	envreader(base)

	if endpointENV {
		BaseURL = os.Getenv(readENV)
	}

	return nil
}

/*AuthenticationAPI is function to get access token*/
func AuthenticationAPI(account string) error {
	var jsonResponse interface{}

	number := regexp.MustCompile(helper.RegexInt()).FindString(BaseURL)
	envLogin := strings.ToUpper(account)
	username := os.Getenv(envLogin + "_USERNAME")
	password := os.Getenv(envLogin + "_PASSWORD")
	readURL := BaseURL + os.Getenv("AUTH_ENDPOINT")

	body := []byte(
		`{
			"grant_type": "` + os.Getenv("GRANT_TYPE") + `",
			"username": "` + username + `",
			"password": "` + password + `",
			"client_id": "` + os.Getenv("API_CLIENT_ID"+"_"+number) + `",
			"client_secret": "` + os.Getenv("API_CLIENT_SECRET"+"_"+number) + `",
			"scope": "` + os.Getenv("SCOPE") + `",
			"` + os.Getenv("COMPANY_ID") + `": "` + os.Getenv("IDENTITY") + `"
		}`)

	client := &http.Client{}
	sendRequest, _ := http.NewRequest("POST", readURL, bytes.NewBuffer(body))

	sendRequest.Header.Set("Content-Type", os.Getenv("CONTENT_TYPE"))
	sendRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	if HTTPResponse, err = client.Do(sendRequest); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	ResponseBody, _ := ioutil.ReadAll(HTTPResponse.Body)
	HTTPJson, _ := jsonpath.Prepare(os.Getenv("JSON_PATH"))

	if err := json.Unmarshal(ResponseBody, &jsonResponse); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	AccessToken, _ = HTTPJson(jsonResponse)

	return nil
}

/*RequestAPI is function to send request*/
func RequestAPI(verbose string, endpoint string, body string) error {
	var stringBody string

	envreader(endpoint)

	if endpointENV {
		endpoint = os.Getenv(readENV)
	}

	readVerbose := strings.ToUpper(verbose)
	requestBody := []byte(body)
	stringBody = string(requestBody)
	regexENV := regexp.MustCompile(helper.RegexReadENV())
	findENV := regexENV.FindAllString(string(requestBody), -1)

	for _, env := range findENV {
		readENV = strings.TrimPrefix(env, "ENV:")
		replaceENV := strings.ReplaceAll(stringBody, env, os.Getenv(readENV))
		stringBody = replaceENV
	}

	readURL := BaseURL + endpoint
	client := &http.Client{}
	sendRequest, _ := http.NewRequest(readVerbose, readURL, bytes.NewBuffer([]byte(stringBody)))

	if AccessToken != nil {
		sendRequest.Header.Add("Authorization", "Bearer "+AccessToken.(string))
	}

	sendRequest.Header.Set("Content-Type", os.Getenv("CONTENT_TYPE"))
	sendRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))

	if HTTPResponse, err = client.Do(sendRequest); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	ResponseBody, _ = ioutil.ReadAll(HTTPResponse.Body)

	return nil
}

/*ResponseStatusAPI is function to get response code*/
func ResponseStatusAPI(response int) error {
	actualCode := HTTPResponse.StatusCode

	if expectCode := (response); actualCode != expectCode {
		log.Panicln(fmt.Errorf("REASON: %s", strconv.Itoa(actualCode)))
	}

	return nil
}
