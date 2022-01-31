package main_test

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/yalp/jsonpath"

	"github.com/golang-automation-v1/features/helper"
	"github.com/golang-automation-v1/features/helper/data"
	"github.com/golang-automation-v1/features/helper/scenario/context"
	"github.com/golang-automation-v1/features/helper/scenario/describe"
	"github.com/golang-automation-v1/features/helper/scenario/it"
	"github.com/golang-automation-v1/features/helper/scenario/specs"
	apisupport "github.com/golang-automation-v1/features/supports/base/api"
)

func TestGolangAutomation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, specs.SUITE)
}

var jsonResponse interface{}

var _ = Describe(describe.ENDPOINT, Ordered, func() {
	BeforeAll(func() {
		env := godotenv.Load("../../../../../.env")
		helper.LogPanicln(env)
	})

	Context(context.EndpointMethod(data.GET), func() {
		It(it.ResponseCodeAssert(strconv.Itoa(data.CODE_200)), func() {
			apisupport.ClientDo(apisupport.SendRequest(data.GET, os.Getenv("ENDPOINT_AUTOMATION"), nil))
			apisupport.DeferResponseReader(false)

			status, _ := jsonpath.Prepare("$.status")
			method, _ := jsonpath.Prepare("$.method")
			endpoint, _ := jsonpath.Prepare("$.endpoint")

			json.Unmarshal(apisupport.ResponseBody, &jsonResponse)

			Expect(apisupport.Response.StatusCode).To(Equal(data.CODE_200))
			Expect(status(jsonResponse)).To(Equal("success"))
			Expect(method(jsonResponse)).To(Equal(data.GET))
			Expect(endpoint(jsonResponse)).To(Equal(os.Getenv("ENDPOINT_AUTOMATION")))
		})
	})

	Context(context.EndpointMethod("POST"), func() {
		It(it.ResponseCodeAssert(strconv.Itoa(data.CODE_201)), func() {
			apisupport.ClientDo(apisupport.SendRequest(data.POST, os.Getenv("ENDPOINT_AUTOMATION"), nil))
			apisupport.DeferResponseReader(false)

			status, _ := jsonpath.Prepare("$.status")
			method, _ := jsonpath.Prepare("$.method")
			endpoint, _ := jsonpath.Prepare("$.endpoint")

			json.Unmarshal(apisupport.ResponseBody, &jsonResponse)

			Expect(apisupport.Response.StatusCode).To(Equal(data.CODE_201))
			Expect(status(jsonResponse)).To(Equal("success"))
			Expect(method(jsonResponse)).To(Equal(data.POST))
			Expect(endpoint(jsonResponse)).To(Equal(os.Getenv("ENDPOINT_AUTOMATION")))
		})
	})

	Context(context.EndpointMethod("PATCH"), func() {
		It(it.ResponseCodeAssert(strconv.Itoa(data.CODE_200)), func() {
			apisupport.ClientDo(apisupport.SendRequest(data.PATCH, os.Getenv("ENDPOINT_AUTOMATION"), nil))
			apisupport.DeferResponseReader(false)

			status, _ := jsonpath.Prepare("$.status")
			method, _ := jsonpath.Prepare("$.method")
			endpoint, _ := jsonpath.Prepare("$.endpoint")

			json.Unmarshal(apisupport.ResponseBody, &jsonResponse)

			Expect(apisupport.Response.StatusCode).To(Equal(data.CODE_200))
			Expect(status(jsonResponse)).To(Equal("success"))
			Expect(method(jsonResponse)).To(Equal(data.PATCH))
			Expect(endpoint(jsonResponse)).To(Equal(os.Getenv("ENDPOINT_AUTOMATION")))
		})
	})

	Context(context.EndpointMethod("PUT"), func() {
		It(it.ResponseCodeAssert(strconv.Itoa(data.CODE_200)), func() {
			apisupport.ClientDo(apisupport.SendRequest(data.PUT, os.Getenv("ENDPOINT_AUTOMATION"), nil))
			apisupport.DeferResponseReader(false)

			status, _ := jsonpath.Prepare("$.status")
			method, _ := jsonpath.Prepare("$.method")
			endpoint, _ := jsonpath.Prepare("$.endpoint")

			json.Unmarshal(apisupport.ResponseBody, &jsonResponse)

			Expect(apisupport.Response.StatusCode).To(Equal(data.CODE_200))
			Expect(status(jsonResponse)).To(Equal("success"))
			Expect(method(jsonResponse)).To(Equal(data.PUT))
			Expect(endpoint(jsonResponse)).To(Equal(os.Getenv("ENDPOINT_AUTOMATION")))
		})
	})

	Context(context.EndpointMethod("DELETE"), func() {
		It(it.ResponseCodeAssert(strconv.Itoa(data.CODE_200)), func() {
			apisupport.ClientDo(apisupport.SendRequest(data.DELETE, os.Getenv("ENDPOINT_AUTOMATION"), nil))
			apisupport.DeferResponseReader(false)

			status, _ := jsonpath.Prepare("$.status")
			method, _ := jsonpath.Prepare("$.method")
			endpoint, _ := jsonpath.Prepare("$.endpoint")

			json.Unmarshal(apisupport.ResponseBody, &jsonResponse)

			Expect(apisupport.Response.StatusCode).To(Equal(data.CODE_200))
			Expect(status(jsonResponse)).To(Equal("success"))
			Expect(method(jsonResponse)).To(Equal(data.DELETE))
			Expect(endpoint(jsonResponse)).To(Equal(os.Getenv("ENDPOINT_AUTOMATION")))
		})
	})
})
