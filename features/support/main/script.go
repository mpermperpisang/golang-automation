//usr/bin/env go run "$0" "$@"; exit

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-automation/features/helper"
	"github.com/golang-automation/features/helper/messages"
	"github.com/golang-automation/features/support"
	"github.com/yalp/jsonpath"
)

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
	helper.LogPanicln(err)

	json.Unmarshal(jsonFile, &reports)
	saveStatus()

	return nil
}

func saveStatus() error {
	for i := 0; i < len(reports); i++ {
		if len(reports[i].Elements) != 0 {
			for s := 0; s < len(reports[i].Elements); s++ {
				arrayStatus = append(arrayStatus, reports[i].Elements[s].Steps[0].Result.Status)
			}
		}
	}

	statusRunCheck(arrayStatus)

	return nil
}

func statusRunCheck(arrayStatus []string) string {
	if helper.Contains(arrayStatus, "failed") {
		statusRun = "FAILED%20:red_circle:"
	} else if helper.Contains(arrayStatus, "passed") {
		statusRun = "SUCCESS%20:green_heart:"
	} else {
		statusRun = outputReplace(messages.NotDetected()) + "%20:no_entry_sign:"
	}

	return statusRun
}

func successPercentageCheck() int {
	var countPassed int

	for i := 0; i < len(arrayStatus); i++ {
		if arrayStatus[i] == "passed" {
			countPassed++
		}
	}

	return countPassed
}

func getGodogInfo() {
	response, err := http.Get("http://localhost:8383/godog-support")
	helper.LogPanicln(err)
	ResponseBody, _ := ioutil.ReadAll(response.Body)

	json.Unmarshal(ResponseBody, &jsonResponse)
}

func getFeatureResponse() string {
	respFeature, err := jsonpath.Read(jsonResponse, "$..feature_tags")
	helper.LogPanicln(err)
	replacer := strings.NewReplacer("[", "", "]", "", " ", "+")
	output := replacer.Replace(fmt.Sprintf("%v", respFeature))

	return output
}

func getPlatformResponse() string {
	respPlatform, err := jsonpath.Read(jsonResponse, "$..platform_name")
	helper.LogPanicln(err)
	replacer := strings.NewReplacer("[", "", "]", "", "-", "+", " ", "%2C+")
	output := replacer.Replace(fmt.Sprintf("%v", respPlatform))

	return output
}

func getDirectoryResponse() string {
	respDirectory, err := jsonpath.Read(jsonResponse, "$..directory")
	helper.LogPanicln(err)
	replacer := strings.NewReplacer("[", "", "]", "", "-", "+", " ", "+%26+")
	output := replacer.Replace(fmt.Sprintf("%v", respDirectory))

	return output
}

func textFormat() string {
	text = "%2AAutomation%20Run%20Result%2A%0D" +
		"%0DStatus%20:%20" + statusRun +
		"%0DTest%20Execution%20tag%20:%20" + featureResponseCheck() +
		"%0DPlatform%20:%20" + platformResponseCheck() +
		"%0DSuccess%20rate%20:%20" + positiveSuccessRateCheck() +
		"%0DHTML%20Report%20:%20" + directoryResponseCheck()

	return text
}

func outputReplace(text string) string {
	replacer := strings.NewReplacer(" ", "%20")
	output := replacer.Replace(fmt.Sprintf("%v", text))

	return output
}

func featureResponseCheck() string {
	if getFeatureResponse() == "" {
		return outputReplace(messages.NotDetected())
	}

	return getFeatureResponse()
}

func platformResponseCheck() string {
	if getPlatformResponse() == "" {
		return outputReplace(messages.NotDetected())
	}

	return getPlatformResponse()
}

func directoryResponseCheck() string {
	if getDirectoryResponse() == "" {
		return outputReplace(messages.NotDetected())
	}

	return "file://" + getDirectoryResponse() + "/test/report/cucumber_report.html"
}

func positiveSuccessRateCheck() string {
	successPercentage := (float64(successPercentageCheck()) / float64(len(arrayStatus))) * 100
	numberIsNegative := math.Signbit(successPercentage)

	if numberIsNegative {
		return outputReplace(messages.NotValid("percentage"))
	}

	return strconv.Itoa(int(successPercentage)) + "%25"
}

func sendNotifTo(apps string) error {
	getGodogInfo()

	resp, err := http.Post("http://localhost:8282/send-"+apps+"?text="+textFormat(), "", nil)
	helper.LogPanicln(err)

	defer resp.Body.Close()

	return nil
}
