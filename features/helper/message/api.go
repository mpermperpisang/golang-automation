package message

import (
	"strconv"
)

/*ResponseCode : pesan untuk response code*/
func ResponseCode(code int) string {
	return "Actual status code : " + strconv.Itoa(code)
}

/*NotMatchDataType : pesan untuk tipe data*/
func NotMatchDataType(data string) string {
	return "Actual data type : " + data
}

/*NotMatchValue : pesan untuk value jsonpath*/
func NotMatchValue(value interface{}) string {
	return "Actual json key value : " + value.(string)
}
