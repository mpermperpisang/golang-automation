package apihelper

import (
	"os"
	"regexp"
	"strings"

	"github.com/golang-automation-v1/features/helper"

	"github.com/cucumber/godog"
)

func EnvReader(body *godog.DocString) string {
	payload := string([]byte(body.Content))
	findENV := regexp.MustCompile(helper.RegexENV()).FindAllString(payload, -1)

	for _, env := range findENV {
		replaceENV := strings.ReplaceAll(payload, env, os.Getenv(helper.TrimENVPrefix(env)))
		payload = replaceENV
	}

	return payload
}
