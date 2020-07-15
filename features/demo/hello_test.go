package demo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type useMock struct {
	mock.Mock
}

func (mock *useMock) helloName(user string) int {
	args := mock.Called(user)

	return args.Int(0)
}

func TestMockHello(t *testing.T) {
	mock := new(useMock)

	mock.On("helloName", "Mper").Return(1)

	m := myStruct{mock}

	assert.Equal(t, m.Hello("Mper"), "Hello Mper!")
}
