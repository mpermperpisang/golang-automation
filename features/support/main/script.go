//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/errors"
	"github.com/golang-automation/features/support"
	"github.com/yalp/jsonpath"
)

var successPercentage int
var text, statusRun string
var arrayStatus []string
var jsonResponse interface{}
var reports support.CucumberReport

func main() {
	statusScenario()
	successPercentageCheck()
	sendNotifTo("slack")
}

func statusScenario() error {
	filename, _ := filepath.Abs("./test/report/cucumber_report.json")
	jsonFile, err := ioutil.ReadFile(filename)
	errors.LogPanicln(err)

	json.Unmarshal(jsonFile, &reports)

	for i := 0; i < len(reports[0].Elements); i++ {
		arrayStatus = append(arrayStatus, reports[0].Elements[i].Steps[0].Result.Status)
	}

	statusRunCheck(arrayStatus)

	return nil
}

func statusRunCheck(arrayStatus []string) string {
	if helper.Contains(arrayStatus, "failed") {
		statusRun = "FAILED"
	} else {
		statusRun = "SUCCESS"
	}

	return statusRun
}

func successPercentageCheck() int {
	var countPassed int

	for i := 0; i < len(reports[0].Elements); i++ {
		if arrayStatus[i] == "passed" {
			countPassed++
		}
	}

	return countPassed
}

func getGodogInfo() {
	response, err := http.Get("http://localhost:8383/godog-support")
	errors.LogPanicln(err)
	ResponseBody, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(ResponseBody, &jsonResponse)
}

func getFeatureResponse() string {
	respFeature, err := jsonpath.Read(jsonResponse, "$..feature_tags")
	errors.LogPanicln(err)
	replacer := strings.NewReplacer("[", "", "]", "", " ", "+%26+")
	output := replacer.Replace(fmt.Sprintf("%v", respFeature))

	return output
}

func getPlatformResponse() string {
	respPlatform, err := jsonpath.Read(jsonResponse, "$..platform_name")
	errors.LogPanicln(err)
	replacer := strings.NewReplacer("[", "", "]", "", "-", "+", " ", "%2C+")
	output := replacer.Replace(fmt.Sprintf("%v", respPlatform))

	return output
}

func getDirectoryResponse() string {
	respDirectory, err := jsonpath.Read(jsonResponse, "$..directory")
	errors.LogPanicln(err)
	replacer := strings.NewReplacer("[", "", "]", "", "-", "+", " ", "+%26+")
	output := replacer.Replace(fmt.Sprintf("%v", respDirectory))

	return output
}

func textFormat() string {
	getGodogInfo()

	successPercentage = int((float64(successPercentageCheck()) / float64(len(reports[0].Elements))) * 100)
	text = "%2AAutomation%20Run%20Result%2A%0D" +
		"%0DStatus%20:%20" + statusRun +
		"%0DTest%20Execution%20tag%20:%20" + getFeatureResponse() +
		"%0DPlatform%20:%20" + getPlatformResponse() +
		"%0DSuccess%20rate%20:%20" + strconv.Itoa(successPercentage) + "%25" +
		"%0DReport%20:%20file://" + getDirectoryResponse() + "/test/report/cucumber_report.html"

	return text
}

func sendNotifTo(apps string) error {
	resp, err := http.Post("http://localhost:8282/send-"+apps+"?text="+textFormat(), "", nil)
	errors.LogPanicln(err)

	defer resp.Body.Close()

	return nil
}
