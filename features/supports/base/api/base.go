package base

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/golang-automation-v1/features/helper"
	"github.com/golang-automation-v1/features/helper/data"
)

var (
	Response     *http.Response
	ResponseBody []byte
	err error
)

func ClientDo(request *http.Request) {
	client := &http.Client{}
	Response, err = client.Do(request)
	helper.LogPanicln(err)
}

func DeferResponseReader(isDefer bool) {
	ResponseBody, err = ioutil.ReadAll(Response.Body)
	helper.LogPanicln(err)

	if (isDefer) {
		defer Response.Body.Close()
	}
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
