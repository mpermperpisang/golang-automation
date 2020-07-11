package stepdefinitions

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/golang-automation/features/demo"
	"github.com/golang-automation/features/helper"
	api "github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/data"
	"github.com/golang-automation/features/helper/messages"
	"github.com/golang-automation/features/support"
	"gopkg.in/yaml.v2"
)

// Example : yaml struct example
type Example struct {
	Description string
	Fruits      map[string][]string
	Methods     map[string]string
}

var usersName, meetName string
var number int
var config Example
var yamlFile []byte

// OpenDWEB : to initiate dweb scenario
func OpenDWEB() error {
	support.DesktopPage.GoToURL(os.Getenv("URL"))
	helper.WaitForLoadingPage(25)

	return nil
}

// OpenMWEB : initiate mweb scenario
func OpenMWEB() error {
	support.MobilePage.GoToURL(os.Getenv("URL"))
	helper.WaitForLoadingPage(25)

	return nil
}

// OpenAndroid : initiate android scenario
func OpenAndroid() error {
	support.AndroidOpenDevice()

	return nil
}

// OpenIOS : initiate ios scenario
func OpenIOS() error {
	support.IOSOpenDevice()

	return nil
}

// BaseAPI : initiate base url for API
func BaseAPI(base string) error {
	api.BaseAPI(base)

	return nil
}

// ResponseStatusAPI : validate response code API
func ResponseStatusAPI(response int) error {
	api.ResponseStatusAPI(response)

	return nil
}

// GivenUserName : assign name to user
func GivenUserName(name string) error {
	usersName = name

	return nil
}

// MeetUserName : call unit
func MeetUserName() error {
	meetName = demo.Hello(usersName)

	return nil
}

// SayHelloName : validate unit
func SayHelloName(greet string) error {
	helper.AssertEqual(greet, meetName, messages.UnitError())

	return nil
}

// GivenData : given some data for example
func GivenData() error {
	number = 5

	return nil
}

// SetData : set some data for example
func SetData() error {
	data.SetDataID(number)

	return nil
}

// GetData : get some data for example
func GetData() error {
	fmt.Println(data.GetDataID())

	return nil
}

// GivenFile : given yaml file
func GivenFile() error {
	var err error

	filename, _ := filepath.Abs("./features/helper/yaml/example.yaml")
	yamlFile, err = ioutil.ReadFile(filename)
	helper.LogPanicln(err)

	return nil
}

// ReadFile : read yaml file
func ReadFile() error {
	err := yaml.Unmarshal(yamlFile, &config)
	helper.LogPanicln(err)

	return nil
}

// PrintContents : print yaml content
func PrintContents() error {
	_, exists := config.Fruits["apple"]
	color := config.Fruits["apple"][0]
	taste := config.Fruits["apple"][1]

	fmt.Printf("Description: %#v\n", config.Description)
	fmt.Printf("Apple fruit in map: %t, color: %v, taste: %v \n", exists, color, taste)

	return nil
}

// GetFunction : call function mapping
func GetFunction() error {
	GivenFile()
	ReadFile()

	exists := config.Methods["f"]
	f := reflect.ValueOf(Example{}).MethodByName(exists).Interface().(func())

	// call dynamic function
	f()

	return nil
}

// Foo : example method
func (Example) Foo() {
	fmt.Println("Call function : foo")
}

// Bar : example method
func (Example) Bar() {
	fmt.Println("Call function : bar")
}
