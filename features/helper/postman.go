package helper

import (
	"os"

	postman "github.com/rbretecher/go-postman-collection"
)

// PostmanCollection : func for using postman collection
func PostmanCollection(jsonFile string) (*postman.Collection, error) {
	file, err := os.Open(jsonFile)
	LogPanicln(err)

	defer file.Close()

	return postman.ParseCollection(file)
}
