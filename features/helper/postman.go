package helper

import (
	"os"

	"github.com/golang-automation/features/helper"
	postman "github.com/rbretecher/go-postman-collection"
)

func postmanCollection(jsonFile string) {
	file, err := os.Open(jsonFile)
	helper.LogPanicln(err)

	defer file.Close()
	parseCollection(file)
}

func parseCollection(file *os.File) *postman.Collection {
	collection, err := postman.ParseCollection(file)
	helper.LogPanicln(err)

	return collection
}
