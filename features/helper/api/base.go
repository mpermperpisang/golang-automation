package apihelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/message"
	"github.com/yalp/jsonpath"
)

/*HTTPResponse variable*/
var HTTPResponse *http.Response

/*BaseURL variable*/
var BaseURL, readENV, oauthToken string
var username, password, number string

/*ResponseBody variable*/
var ResponseBody []byte
var endpointENV bool

/*AccessToken variable*/
var AccessToken interface{}
var err error

func envReader(env string) error {
	endpointENV = strings.HasPrefix(env, "ENV:")
	readENV = strings.TrimPrefix(env, "ENV:")

	return nil
}

/*BaseAPI is function to set base url for API*/
func BaseAPI(base string) error {
	envReader(base)

	if endpointENV {
		BaseURL = os.Getenv(readENV)
	}

	return nil
}

func oauthURL() error {
	oauthToken = BaseURL + os.Getenv("AUTH_ENDPOINT")

	return nil
}

func varLogin(account string) error {
	envLogin := strings.ToUpper(account)
	username = os.Getenv(envLogin + "_USERNAME")
	password = os.Getenv(envLogin + "_PASSWORD")

	return nil
}

func stagingNumber() error {
	number = regexp.MustCompile(helper.RegexInt()).FindString(BaseURL)

	if number == "" {
		number = regexp.MustCompile(helper.RegexBaseURL()).FindString(BaseURL)
	}

	return nil
}

func clientResponse(sendRequest *http.Request) error {
	client := &http.Client{}

	if HTTPResponse, err = client.Do(sendRequest); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	ResponseBody, _ = ioutil.ReadAll(HTTPResponse.Body)

	return nil
}

/*AuthenticationAPI is function to get access token*/
func AuthenticationAPI(account string) error {
	var jsonResponse interface{}

	oauthURL()
	varLogin(account)
	stagingNumber()

	number := strings.ToUpper(number)
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
	sendRequest, _ := http.NewRequest("POST", oauthToken, bytes.NewBuffer(body))
	HTTPJson, _ := jsonpath.Prepare(os.Getenv("JSON_PATH"))

	sendRequest.Header.Set("Content-Type", os.Getenv("CONTENT_TYPE"))
	sendRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))
	clientResponse(sendRequest)

	if err := json.Unmarshal(ResponseBody, &jsonResponse); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	AccessToken, _ = HTTPJson(jsonResponse)

	return nil
}

/*RequestAPI is function to send request*/
func RequestAPI(verbose string, endpoint string, body string) error {
	var stringBody string

	envReader(endpoint)

	if endpointENV {
		endpoint = os.Getenv(readENV)
	}

	readVerbose := strings.ToUpper(verbose)
	requestBody := []byte(body)
	stringBody = string(requestBody)
	regexENV := regexp.MustCompile(helper.RegexReadENV())
	findENV := regexENV.FindAllString(string(requestBody), -1)

	for _, env := range findENV {
		envReader(env)
		replaceENV := strings.ReplaceAll(stringBody, env, os.Getenv(readENV))
		stringBody = replaceENV
	}

	readURL := BaseURL + endpoint
	sendRequest, _ := http.NewRequest(readVerbose, readURL, bytes.NewBuffer([]byte(stringBody)))

	if AccessToken != nil {
		sendRequest.Header.Add("Authorization", "Bearer "+AccessToken.(string))
	}

	sendRequest.Header.Set("Content-Type", os.Getenv("CONTENT_TYPE"))
	sendRequest.Header.Set("User-Agent", os.Getenv("USER_AGENT"))
	clientResponse(sendRequest)

	return nil
}

/*ResponseStatusAPI is function to get response code*/
func ResponseStatusAPI(response int) error {
	actualCode := HTTPResponse.StatusCode

	helper.AssertEqual(response, actualCode, message.ResponseCode(actualCode))

	return nil
}
