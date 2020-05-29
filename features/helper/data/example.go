package data

// ExampleData : data for example
type ExampleData struct {
	id int
}

var data ExampleData
var id = data.id

/*SetDataID : set data for ID*/
func SetDataID(id int) {
	data.id = id
}

/*GetDataID : get data for ID*/
func GetDataID() int {
	return id
}
