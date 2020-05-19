package data

// ExampleData :
type ExampleData struct {
	id int
}

var data ExampleData

/*SetDataID : set data for ID*/
func SetDataID(id int) {
	data.id = id
}

/*GetDataID : get data for ID*/
func GetDataID() int {
	return data.id
}
