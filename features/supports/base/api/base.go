package base

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"golang-automation/features/helper"
	"golang-automation/features/helper/data"
)

var (
	Response     *http.Response
	ResponseBody []byte
)

func ClientDo(request *http.Request) {
	var err error

	client := &http.Client{}
	Response, err = client.Do(request)
	helper.LogPanicln(err)
	ResponseBody, err = ioutil.ReadAll(Response.Body)
	helper.LogPanicln(err)

	defer Response.Body.Close()
}

func SendRequest(method, url string, body io.Reader) *http.Request {
	if helper.HasENVPrefix(url) {
		url = os.Getenv(helper.TrimENVPrefix(url))
	}

	endpoint := os.Getenv(data.BASE_API) + url
	request, err := http.NewRequest(method, endpoint, body)
	helper.LogPanicln(err)

	return request
}
