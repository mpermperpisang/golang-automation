package data

type exampleData struct {
	id int
}

var data exampleData
var id = data.id

// SetDataID : set data for ID
func SetDataID(id int) {
	data.id = id
}

// GetDataID : get data for ID
func GetDataID() int {
	return id
}
