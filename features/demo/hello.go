package demo

import "fmt"

type myMock interface {
	helloName(user string) int
}

type myStruct struct {
	mock myMock
}

// Hello : say hello to user with mock
func (m myStruct) Hello(user string) string {
	if m.mock.helloName(user) == 0 {
		return "Hello Dude!"
	}

	return fmt.Sprintf("Hello %v!", user)
}

// HelloUser : say hello to user without mock
func HelloUser(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}

	return fmt.Sprintf("Hello %v!", user)
}
