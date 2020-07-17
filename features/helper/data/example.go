package data

type exampleData struct {
	id int
}

var data exampleData

// SetID : set data for ID
func SetID(id int) {
	data.id = id
}

// ID : get data for ID
func ID() int {
	return data.id
}
